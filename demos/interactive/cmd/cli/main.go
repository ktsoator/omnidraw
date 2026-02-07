package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type Prize struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

const (
	symbolSelect = "●"
	symbolCheck  = "✓"
	symbolCross  = "✗"
)

var (
	styleSubtle  = color.New(color.FgHiBlack)
	styleAccent  = color.New(color.FgHiCyan, color.Bold)
	styleSuccess = color.New(color.FgGreen)
	styleError   = color.New(color.FgRed)

	serverURL string
	client    *http.Client
)

func main() {
	flag.StringVar(&serverURL, "server", "http://localhost:8080/prize", "Server API URL")
	flag.Parse()

	client = &http.Client{Timeout: 10 * time.Second}

	fmt.Println()
	styleSubtle.Print("  OMNIDRAW ")
	styleAccent.Println("INTERACTIVE")
	fmt.Println()

	items := []string{"Import Prizes", "Manual Upload", "Exit"}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   fmt.Sprintf("  {{ \"%s\" | cyan }} {{ . | bold }}", symbolSelect),
		Inactive: "    {{ . | faint }}",
		Selected: fmt.Sprintf("  {{ \"%s\" | green }} {{ . | faint }}", symbolCheck),
	}

	prompt := promptui.Select{
		Label:        "",
		Items:        items,
		Templates:    templates,
		Size:         4,
		HideSelected: true,
	}

	for {
		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		switch index {
		case 0:
			importPrizes()
		case 1:
			uploadPrize()
		case 2:
			fmt.Println()
			styleAccent.Printf("  %s ", symbolCheck)
			styleSubtle.Println("Goodbye.")
			fmt.Println()
			return
		}
		fmt.Println()
	}
}

func importPrizes() {
	prompt := promptui.Prompt{
		Label: "  File Path",
		Templates: &promptui.PromptTemplates{
			Prompt: "  {{ . | cyan }}: ",
		},
		Default: "test_prizes.json",
	}

	filename, err := prompt.Run()
	if err != nil {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		styleError.Printf("  %s Failed to open file: %v\n", symbolCross, err)
		return
	}
	defer file.Close()

	var prizes []Prize
	if err := json.NewDecoder(file).Decode(&prizes); err != nil {
		styleError.Printf("  %s Failed to parse JSON: %v\n", symbolCross, err)
		return
	}

	doUpload(prizes)
}

func doUpload(prizes []Prize) {
	if len(prizes) == 0 {
		styleSubtle.Println("  No prizes to upload.")
		return
	}

	payload := map[string][]Prize{"prizes": prizes}
	body, _ := json.Marshal(payload)

	resp, err := client.Post(serverURL+"/upload", "application/json", bytes.NewBuffer(body))
	if err != nil {
		styleError.Printf("  %s Network error: %v\n", symbolCross, err)
		return
	}
	defer resp.Body.Close()

	var res Result
	json.NewDecoder(resp.Body).Decode(&res)

	if res.Code == 0 {
		styleSuccess.Printf("  %s Successfully uploaded %d prizes\n", symbolCheck, len(prizes))
	} else {
		styleError.Printf("  %s Upload failed: %s\n", symbolCross, res.Msg)
	}
}

func uploadPrize() {
	fmt.Println()
	styleSubtle.Println("  Enter prize details (leave name empty to finish):")

	var prizes []Prize
	for {
		namePrompt := promptui.Prompt{
			Label: "    Prize Name",
			Templates: &promptui.PromptTemplates{
				Prompt: "  {{ . }}: ",
			},
		}
		name, _ := namePrompt.Run()
		name = strings.TrimSpace(name)
		if name == "" {
			break
		}

		countPrompt := promptui.Prompt{
			Label: "    Count",
			Templates: &promptui.PromptTemplates{
				Prompt: "  {{ . }}: ",
			},
			Default: "1",
		}
		countStr, _ := countPrompt.Run()
		var count int
		fmt.Sscanf(countStr, "%d", &count)

		prizes = append(prizes, Prize{Name: name, Count: count})
	}

	doUpload(prizes)
}

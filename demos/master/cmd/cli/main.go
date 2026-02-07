package main

import (
	"bufio"
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

const baseURL = "http://localhost:8080/master"

type Result struct {
	Message   string `json:"message"`
	Winner    string `json:"winner,omitempty"`
	Remaining int    `json:"remaining,omitempty"`
}

const (
	symbolPrompt = "❯"
	symbolSelect = "●"
	symbolCheck  = "✓"
	symbolCross  = "✗"
	symbolTrophy = "★"
)

var (
	// Muted, sophisticated palette
	styleSubtle    = color.New(color.FgHiBlack)
	styleAccent    = color.New(color.FgHiCyan, color.Bold)
	styleNormal    = color.New(color.FgWhite)
	styleSuccess   = color.New(color.FgGreen)
	styleError     = color.New(color.FgRed)
	styleHighlight = color.New(color.FgHiWhite, color.Bold)

	serverURL string
	client    *http.Client
)

func main() {
	flag.StringVar(&serverURL, "server", "http://localhost:8080/master", "Server API URL")
	flag.Parse()

	client = &http.Client{Timeout: 10 * time.Second}

	// Clean header with just spacing
	fmt.Println()
	styleSubtle.Print("  OMNIDRAW ")
	styleAccent.Println("MASTER")
	fmt.Println()

	items := []string{"Import List", "Draw Winner", "Exit"}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   fmt.Sprintf("  {{ \"%s\" | cyan }} {{ . | bold }}", symbolSelect),
		Inactive: "    {{ . | faint }}",
		Selected: fmt.Sprintf("  {{ \"%s\" | green }} {{ . | faint }}", symbolCheck),
		Details:  "",
	}

	prompt := promptui.Select{
		Label:        "", // No label for cleaner look
		Items:        items,
		Templates:    templates,
		Size:         4,
		HideSelected: true, // Keep it clean
	}

	for {
		index, _, err := prompt.Run()
		if err != nil {
			return
		}

		switch index {
		case 0:
			importPlayers()
		case 1:
			drawWinner()
		case 2:
			fmt.Println()
			styleAccent.Printf("  %s ", symbolCheck)
			styleSubtle.Println("Session ended. Goodbye.")
			fmt.Println()
			return
		}
		fmt.Println()
	}
}

func importPlayers() {
	prompt := promptui.Prompt{
		Label: "File",
		Templates: &promptui.PromptTemplates{
			Prompt:  fmt.Sprintf("  {{ \"%s\" | cyan }} {{ . | faint }}: ", symbolPrompt),
			Valid:   fmt.Sprintf("  {{ \"%s\" | green }} {{ . | faint }}: ", symbolPrompt),
			Invalid: fmt.Sprintf("  {{ \"%s\" | red }} {{ . | faint }}: ", symbolCross),
			Success: fmt.Sprintf("  {{ \"%s\" | green }} File: ", symbolCheck),
		},
		Default:     "test_players.txt",
		HideEntered: true,
	}

	filename, err := prompt.Run()
	if err != nil {
		return
	}
	filename = strings.TrimSpace(filename)

	file, err := os.Open(filename)
	if err != nil {
		styleError.Printf("  %s %v\n", symbolCross, err)
		return
	}
	defer file.Close()

	// [File reading logic remains the same]
	var players []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if name := strings.TrimSpace(scanner.Text()); name != "" {
			players = append(players, name)
		}
	}

	if len(players) == 0 {
		styleError.Printf("  %s Empty file\n", symbolCross)
		return
	}

	payload := map[string][]string{"players": players}
	body, _ := json.Marshal(payload)

	resp, err := client.Post(serverURL+"/players", "application/json", bytes.NewBuffer(body))
	if err != nil {
		styleError.Printf("  %s Network error\n", symbolCross)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		styleSuccess.Printf("  %s Loaded %d players\n", symbolCheck, len(players))
	} else {
		styleError.Printf("  %s Server rejected import\n", symbolCross)
	}
}

func drawWinner() {
	resp, err := client.Get(serverURL + "/draw")
	if err != nil {
		styleError.Printf("  %s Network error\n", symbolCross)
		return
	}
	defer resp.Body.Close()

	var res Result
	json.NewDecoder(resp.Body).Decode(&res)

	if resp.StatusCode == http.StatusOK {
		if res.Winner != "" {
			// Subtle "drawing" feel
			styleSubtle.Print("  Drawing")
			for i := 0; i < 3; i++ {
				time.Sleep(150 * time.Millisecond)
				styleSubtle.Print(".")
			}
			fmt.Print("\r") // Clear the "Drawing..." line

			// Elegant result display
			fmt.Print("  ")
			styleHighlight.Printf("%s %s\n", symbolTrophy, res.Winner)
			styleSubtle.Printf("    %d left in pool\n", res.Remaining)
		} else {
			styleError.Printf("  %s %s\n", symbolCross, res.Message)
		}
	} else {
		styleError.Printf("  %s Draw failed\n", symbolCross)
	}
}

package service

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

var (
	playerPool []string
	mu         sync.Mutex
)

func SetDrawPlayers(users []string) {
	mu.Lock()
	defer mu.Unlock()

	filtered := make([]string, 0, len(users))
	for _, user := range users {
		if trimmed := strings.TrimSpace(user); trimmed != "" {
			filtered = append(filtered, trimmed)
		}
	}
	playerPool = filtered
}

func DrawPlayer() (string, string) {
	mu.Lock()
	defer mu.Unlock()

	count := len(playerPool)
	if count == 0 {
		return "", "No players remaining for the draw!"
	}

	index := rand.Intn(count)
	winner := playerPool[index]
	playerPool = append(playerPool[:index], playerPool[index+1:]...)

	return winner, fmt.Sprintf("%d", len(playerPool))
}

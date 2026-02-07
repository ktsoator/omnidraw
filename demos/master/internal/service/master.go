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

func SetDrawPlayers(players []string) {
	mu.Lock()
	defer mu.Unlock()

	seen := make(map[string]bool)
	filtered := make([]string, 0, len(players))
	for _, p := range players {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" && !seen[trimmed] {
			filtered = append(filtered, trimmed)
			seen[trimmed] = true
		}
	}
	playerPool = filtered
}

func DrawPlayer() (string, int, error) {
	mu.Lock()
	defer mu.Unlock()

	count := len(playerPool)
	if count == 0 {
		return "", 0, fmt.Errorf("no players remaining for the draw")
	}

	index := rand.Intn(count)
	winner := playerPool[index]

	// O(1) removal: swap with last element and shrink
	playerPool[index] = playerPool[count-1]
	playerPool = playerPool[:count-1]

	return winner, len(playerPool), nil
}

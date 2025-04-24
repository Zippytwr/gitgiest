package parsers

import (
	"fmt"
	"strings"
)

func ParseAuthorStats(output []byte) {
	lines := strings.Split(string(output), "\n")
	stats := make(map[string]int)
	for _, line := range lines {
		if line == "" {
			continue
		}
		stats[line]++
	}
	for author, count := range stats {
		fmt.Printf("\033[1;36m%s\033[0m → \033[1;32m%d commits\033[0m\n", author, count)
	}
}

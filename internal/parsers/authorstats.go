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
	for key, value := range stats {
		if value < 10 {
			fmt.Printf("\033[1;35mЛентяй всего лишь %d коммитов у %s\033[0m\n", value, key)
		}
		if value > 200 {
			fmt.Printf("\033[1;34m%d — Середняк!!! дефолт коммитов - %s\033[0m\n", value, key)
		}
		if value > 1000 {
			fmt.Printf("\033[1;34m%d — Маг кодеееер!!! много коммитов - %s\033[0m!!\n", value, key)
		}
	}

}

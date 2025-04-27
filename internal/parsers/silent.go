package parsers

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func ParseSilent(output []byte) {
	commitHashes := strings.Split(string(output), "\n")
	insertionRegex := regexp.MustCompile(`(\d+) insertions\(\+\)`)

	for _, commitHash := range commitHashes {
		if commitHash == "" {
			continue
		}

		cmdShowCommit := exec.Command("git", "show", "--shortstat", commitHash)
		showOutput, _ := cmdShowCommit.Output()
		showOutputStr := string(showOutput)

		matches := insertionRegex.FindStringSubmatch(showOutputStr)
		if len(matches) < 2 {
			continue
		}

		insertionsCount, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Printf("Ошибка конвертации количества вставок: %v\n", err)
			continue
		}

		switch {
		case insertionsCount == 0:
			fmt.Printf("Лентяй! Ни одной вставки в коммите %s.\n", commitHash)
		case insertionsCount < 5:
			fmt.Printf("Очень мало вставок (%d) в коммите %s.\n", insertionsCount, commitHash)
		case insertionsCount > 50:
			fmt.Printf("Хороший коммит %s - много вставок (%d).\n", commitHash, insertionsCount)
		default:
			fmt.Printf("Нормальное количество вставок (%d) в коммите %s.\n", insertionsCount, commitHash)
		}
	}
}

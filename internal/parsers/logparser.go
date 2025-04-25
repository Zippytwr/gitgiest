package parsers

import (
	"fmt"
	"strings"
)

// Структура для хранения информации о коммитах
type Commit struct {
	Hash    string
	Author  string
	Date    string
	Message string
}

// Парсинг логов git и возврат информации о коммитах
func ParseGitLog(output []byte) []Commit {
	lines := strings.Split(string(output), "\n")
	var commits []Commit
	for _, line := range lines {
		parts := strings.SplitN(line, "|", 4)
		if len(parts) < 4 {
			continue
		}
		commits = append(commits, Commit{
			Hash:    parts[0],
			Author:  parts[1],
			Date:    parts[2],
			Message: parts[3],
		})
	}
	return commits
}

// Вывод коммитов с красивым форматированием
func PrintCommits(commits []Commit) {
	for _, commit := range commits {
		fmt.Printf("Смска: \033[1;34m%s\033[0m\n", commit.Message)
		fmt.Printf("\033[1;33m%s\033[0m нашаманил \033[1;32m%s\033[0m\n", commit.Date, commit.Author)
		fmt.Printf("Если умный: \033[2m%s\033[0m\n\n", commit.Hash[:8])
	}
}

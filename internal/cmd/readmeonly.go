package cmd

import (
	"fmt"
	"gitgeist/internal/git"
	"github.com/urfave/cli/v2"
	"log"
	"os/exec"
	"strings"
)

var readmeonly = &cli.Command{
	Name:  "readme",
	Usage: "Show readme only workers",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "repo",
			Usage: "Path to the git repository",
		},
	},
	Action: func(c *cli.Context) error {
		repoPath := c.String("repo")
		if repoPath == "" {
			repoPath = "./"
		}
		output, err := git.RunGitReadmeOnly(repoPath)
		if err != nil {
			log.Fatal(err)
		}
		outputString := string(output)
		result := strings.Split(outputString, "\n")
		fmt.Println("📝 Коммиты, изменяющие только README.md\n")
		for _, commitHash := range result {
			if commitHash == "" {
				continue
			}
			cmdAuthor := exec.Command("git", "show", "-s", "--format=%an", commitHash)
			authorBytes, err := cmdAuthor.Output()
			author := "Неизвестный"
			if err == nil {
				author = strings.TrimSpace(string(authorBytes))
			}
			cmd := exec.Command("git", "show", "--name-only", "--pretty=", commitHash)
			out, err := cmd.Output()
			if err != nil {
				fmt.Println("Ошибка при выполнении git show:", err)
				continue
			}

			files := strings.Split(strings.TrimSpace(string(out)), "\n")
			if len(files) == 1 && strings.ToLower(files[0]) == "readme.md" {

				fmt.Printf("\033[1;33m⚠️  %s — Обнаружен Лентяй! (%s)\n   🛋️ Звание: README-only Commando\033[0m\n", commitHash, author)
			} else {
				fmt.Printf("\033[1;32m✅ %s — Хорошая работа, %s! Реальные изменения внесены.\033[0m\n", commitHash, author)
			}
		}
		return nil
	},
}

package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var silentCommits = &cli.Command{
	Name:  "silent",
	Usage: "Detects commits that made minimal changes — perfect for spotting \"silent commits\" (tiny edits, minor fixes, or micro-changes).\n",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "repo",
			Usage: "Path to the git repository",
		},
	},
	Action: func(c *cli.Context) error {
		cmdGetCommits := exec.Command("git", "log", "--pretty=format:%H")
		commitHashesRaw, _ := cmdGetCommits.Output()
		commitHashes := strings.Split(string(commitHashesRaw), "\n")

		// Регулярка для поиска количества вставок
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
				fmt.Println("Ошибка конвертации количества вставок:", err)
				continue
			}

			if insertionsCount > 50 {
				fmt.Printf("Дефолт коммит %s тут даже не на что смотреть, но вот тебе сколько вставок: %d \n", commitHash, insertionsCount)
			}
			if insertionsCount < 5 {
				fmt.Printf("Очень мало вставок. Столько: %s", insertionsCount)
			}
			if insertionsCount < 1 {
				fmt.Printf("Лентяй!!! столько вставок: %s, вот его коммит: %d", insertionsCount, commitHashes)
			}
		}

		return nil
	},
}

package cmd

import (
	"gitgeist/internal/git"
	"gitgeist/internal/parsers"
	"github.com/urfave/cli/v2"
	"log"
)

var logCommand = &cli.Command{
	Name:  "log",
	Usage: "Show raw git log",
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

		// Получение логов из git
		output, err := git.RunGitLog(repoPath)
		if err != nil {
			log.Fatal(err)
		}

		// Парсинг логов и вывод коммитов
		commits := parsers.ParseGitLog(output)
		parsers.PrintCommits(commits)
		return nil
	},
}

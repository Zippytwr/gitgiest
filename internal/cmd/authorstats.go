package cmd

import (
	"gitgeist/internal/git"
	"gitgeist/internal/parsers"
	"github.com/urfave/cli/v2"
	"log"
)

var authorstats = &cli.Command{
	Name:  "authorstats",
	Usage: "Show commit counts per author",
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

		// Получение статистики по авторам
		output, err := git.RunGitAuthorStats(repoPath)
		if err != nil {
			log.Fatal(err)
		}

		// Парсинг статистики и вывод результатов
		parsers.ParseAuthorStats(output)
		return nil
	},
}

package main

import (
	"gitgeist/authorstats"
	"gitgeist/git"
	"gitgeist/logparser"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "gitgeist",
		Usage: "Git log analyzer with streak calculation",
		Commands: []*cli.Command{
			{
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
					commits := logparser.ParseGitLog(output)
					logparser.PrintCommits(commits)
					return nil
				},
			},
			{
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
					authorstats.ParseAuthorStats(output)
					return nil
				},
			},
		},
	}

	// Запуск приложения
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package cmd

import (
	"gitgeist/internal/git"
	"gitgeist/internal/parsers"
	"github.com/urfave/cli/v2"
	"log"
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
		parsers.ReadmeParser(output)
		return nil
	},
}

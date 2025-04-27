package cmd

import (
	"gitgeist/internal/git"
	"gitgeist/internal/parsers"
	"github.com/urfave/cli/v2"
	"log"
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
		repoPath := c.String("repo")
		if repoPath == "" {
			repoPath = "./"
		}
		output, err := git.RunGitSilent(repoPath)
		if err != nil {
			log.Fatal(err)
		}
		parsers.ParseSilent(output)
		return nil
	},
}

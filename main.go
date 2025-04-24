package main

import (
	"gitgeist/internal/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:     "gitgeist",
		Usage:    "Git log analyzer with streak calculation",
		Commands: cmd.Commands,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

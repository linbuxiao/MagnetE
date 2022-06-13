package main

import (
	"log"
	"os"

	"github.com/linbuxiao/magnete/cmd/action"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "magnet E",
		Commands: []*cli.Command{
			action.GetCommand,
			action.PingCommand,
			action.YamlCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

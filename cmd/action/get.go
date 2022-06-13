package action

import (
	"fmt"

	"github.com/linbuxiao/magnete/gen/collector"
	"github.com/urfave/cli/v2"
)

var GetCommand = &cli.Command{
	Name: "get",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "engine",
			Aliases: []string{"e"},
		},
		&cli.StringFlag{
			Name:     "query",
			Aliases:  []string{"q"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "format",
			Aliases: []string{"f"},
		},
	},
	Action: get,
}

func get(ctx *cli.Context) error {
	e := ctx.String("engine")
	q := ctx.String("query")
	f := ctx.String("format")
	if e == "" {
		e = "btsow"
	}
	list, err := collector.Get(e, q)
	if err != nil {
		return err
	}
	b, err := formatOutput(f, list)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

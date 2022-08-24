package main

import (
	"log"
	"os"
	"sort"
	"zssh/handle"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{Name: "list", Aliases: []string{"l"}, Usage: "Get List", Action: handle.List},
			{Name: "init", Usage: "init config file, danger", Action: handle.InitAction},
			{Name: "add", Usage: "add [name][pwd][host]", Action: handle.Add},
			{Name: "to", Usage: "to [index]", Action: handle.To},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

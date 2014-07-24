package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/elentok/gever/version"
)

var Show = cli.Command{
	Name:      "show",
	ShortName: "s",
	Usage:     "shows the version of the current project",
	Flags: []cli.Flag{
		cli.BoolFlag{"verbose, v", "Show more details"},
	},
	Action: show,
}

func show(c *cli.Context) {
	v, err := version.Find(".", c.Bool("verbose"))
	if err != nil {
		println("Version not found")
		os.Exit(1)
	}

	fmt.Println(v.ToString())
}

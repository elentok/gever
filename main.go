package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/elentok/gever/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "gever"
	app.Usage = "semver-based versioning tool"
	app.Commands = []cli.Command{
		commands.Show,
		commands.Create,
	}

	app.Run(os.Args)
}

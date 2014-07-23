package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/elentok/gever/commands"
)

// this value is replaced when building:
//
//   go build -ldflags "-X main.version #{version}" main.go
var version string = "DEV"

func main() {
	app := cli.NewApp()
	app.Name = "gever"
	app.Usage = "semver-based versioning tool"
	app.Version = version
	app.Commands = []cli.Command{
		commands.Show,
		commands.Create,
	}

	app.Run(os.Args)
}

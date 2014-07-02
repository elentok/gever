package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/elentok/gever/version"
)

var Create = cli.Command{
	Name:      "create",
	ShortName: "c",
	Usage:     "creates new versions",
	Flags: []cli.Flag{
		cli.BoolFlag{"noop, n", "Only show the new version, don't do anything"},
		cli.BoolFlag{"verbose, v", "Show more information"},
	},
	Subcommands: []cli.Command{
		newCreateCommand("hotfix", "h", createHotfix),
		newCreateCommand("rc", "", createRC),
		newCreateCommand("patch", "p", createPatch),
		newCreateCommand("minor", "mi", createMinor),
		newCreateCommand("major", "ma", createMajor),
	},
}

type versionCreator func(v version.Version)

func newCreateCommand(name, shortName string, create versionCreator) cli.Command {
	return cli.Command{
		Name:      name,
		ShortName: shortName,
		Usage:     fmt.Sprintf("creates a %s", name),
		Action: func(c *cli.Context) {
			v, err := version.Find(".", c.Bool("verbose"))
			if err != nil {
				exitOnError(err)
			}
			create(v)
			println(v.ToString())

			if c.Bool("noop") {
				return
			}

			// TODO: what to do now? (create git tag? update the version file?)
		},
	}

}

func createHotfix(v version.Version) { v.CreateHotfix() }
func createRC(v version.Version)     { v.CreateRC() }
func createPatch(v version.Version)  { v.CreatePatch() }
func createMinor(v version.Version)  { v.CreateMinor() }
func createMajor(v version.Version)  { v.CreateMajor() }

func exitOnError(err error) {
	fmt.Printf("ERROR: %s\n", err)
	os.Exit(1)
}

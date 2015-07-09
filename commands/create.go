package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/elentok/gever/git"
	"github.com/elentok/gever/utils"
	"github.com/elentok/gever/version"
)

var Create = cli.Command{
	Name:      "create",
	ShortName: "c",
	Usage:     "creates new versions",
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
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "noop, n",
				Usage: "Only show the new version, don't do anything"},
			cli.BoolFlag{
				Name:  "quiet, q",
				Usage: "Don't ask before tagging"},
			cli.BoolFlag{
				Name:  "verbose, v",
				Usage: "Show more information"},
		},
		Action: func(c *cli.Context) {
			v, err := version.Find(".", c.Bool("verbose"))
			if err != nil {
				exitOnError(err)
			}
			if v == nil {
				println("Previous version not found, defaulting to 0.0.0")
				v = version.New(0, 0, 0, "")
			}
			create(v)

			if c.Bool("noop") {
				fmt.Println(v.ToString())
				return
			}

			tag := fmt.Sprintf("v%s", v.ToString())
			msg := fmt.Sprintf("Create git tag '%s'", tag)
			if c.Bool("quiet") || utils.Confirm(msg, true) {
				repo, err := git.NewRepo(".")
				if err != nil {
					exitOnError(err)
				}
				msg = fmt.Sprintf("Bumped version to %v", tag)
				fmt.Println(msg)
				repo.Tag(tag, msg)
			}
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

package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "simple. distributed. version control."
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "turn on debug logging",
		},
		cli.StringFlag{
			Name:  "cwd, c",
			Value: ".",
			Usage: "set the working directory to `CWD`",
		},
		cli.BoolFlag{
			Name:  "pretty",
			Usage: "force pretty, human-oriented output",
		},
		cli.BoolFlag{
			Name:  "plain",
			Usage: "force plain output (no terminal colors)",
		},
		cli.BoolFlag{
			Name:  "ugly",
			Usage: "use ugly, machine-parseable output",
		},
	}
	// so we have some consistency
	sort.Sort(cli.FlagsByName(app.Flags))
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "initialize a new repository",
			Action: Init,
		},
		{
			Name:    "status",
			Aliases: []string{"stat"},
			Usage:   "display repository state",
			Action:  Status,
		},
		{
			Name:  "commit",
			Usage: "commit changes to repository",
			Action: func(c *cli.Context) error {
				fmt.Println("would run commit")
				return nil
			},
		},
		{
			Name:  "undo",
			Usage: "rollback changes",
			Action: func(c *cli.Context) error {
				fmt.Println("would run undo")
				return nil
			},
		},
		{
			Name:  "diff",
			Usage: "show differences between snapshots",
			Action: func(c *cli.Context) error {
				fmt.Println("would run diff")
				return nil
			},
		},
		{
			Name:  "log",
			Usage: "log of what happened",
			Action: func(c *cli.Context) error {
				fmt.Println("would run log")
				return nil
			},
		},
	}
	app.Run(os.Args)
}

package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func Status(c *cli.Context) error {
	var (
		// TODO add a 'Before' hook to auto-generate the repo state,
		// handle init errors, etc
		conf = NewConfigWithDefaults()
		fs   = NewDefaultFilesystem()
	)
	repo := NewRepo(conf, fs)
	fmt.Println(repo)
	return nil
}

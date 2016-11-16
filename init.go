package main

import "github.com/urfave/cli"

func Init(c *cli.Context) error {
	var (
		// TODO parse global & local config files
		conf = NewConfigWithDefaults()
		fs   = NewDefaultFilesystem()
	)
	return NewRepo(conf, fs).Commit()
}

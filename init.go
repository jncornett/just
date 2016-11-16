package main

import (
	"github.com/spf13/afero"
	"github.com/urfave/cli"
)

func Init(c *cli.Context) error {
	return NewRepo(NewConfigFromCmdline(c), afero.NewMemMapFs()).Commit()
}

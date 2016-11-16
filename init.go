package main

import (
	"github.com/spf13/afero"
	"github.com/urfave/cli"
)

func Init(c *cli.Context) error {
	return NewRepo(NewConfigWithDefaults(), afero.NewMemMapFs()).Commit()
}

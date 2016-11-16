package main

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/urfave/cli"
)

func Status(c *cli.Context) error {
	repo := NewRepo(NewConfigFromCmdline(c), afero.NewMemMapFs())
	fmt.Println(repo)
	return nil
}

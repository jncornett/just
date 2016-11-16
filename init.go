package main

import (
	"log"

	"github.com/urfave/cli"
)

func Init(c *cli.Context) error {
	repo := NewRepo(NewDefaultFilesystem())
	err := repo.Commit()
	if err != nil {
		log.Print(err)
	}
	return err
}

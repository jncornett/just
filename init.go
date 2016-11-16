package main

import (
	"log"

	"github.com/urfave/cli"
)

func Init(c *cli.Context) error {
	var (
		conf = NewConfigWithDefaults()
		fs   = NewDefaultFilesystem()
		repo = NewRepo(conf, fs)
	)
	err := repo.Commit()
	if err != nil {
		log.Print(err)
	}
	return err
}

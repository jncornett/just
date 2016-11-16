package main

import "github.com/spf13/afero"

type Repo struct {
	Conf *Config
	Fs   afero.Fs
}

func NewRepo(conf *Config, fs afero.Fs) *Repo {
	return &Repo{conf, fs}
}

func (r *Repo) Commit() error {
	return r.Fs.Mkdir(r.Conf.RepoDirname, 0)
}

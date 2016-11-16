package main

const RepoDirname = ".just"

type Repo struct {
	Conf *Config
	Fs   Filesystem
}

func NewRepo(conf *Config, fs Filesystem) *Repo {
	return &Repo{conf, fs}
}

func (r *Repo) Commit() error {
	r.Fs.Mkdir(r.Conf.RepoDirname)
	return nil
}

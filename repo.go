package main

type Repo struct {
	Conf *Config
	Fs   Filesystem
}

func NewRepo(conf *Config, fs Filesystem) *Repo {
	return &Repo{conf, fs}
}

func (r *Repo) Commit() error {
	return r.Fs.Mkdir(r.Conf.RepoDirname)
}

package main

const RepoDirname = ".just"

type Repo struct {
	Fs Filesystem
}

func NewRepo(fs Filesystem) *Repo {
	return &Repo{fs}
}

func (r *Repo) Commit() error {
	r.Fs.Mkdir(".just")
	return nil
}

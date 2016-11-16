package main

const DefaultRepoDirname = ".just"

type Config struct {
	RepoDirname string
}

func NewConfigWithDefaults() *Config {
	return &Config{
		RepoDirname: DefaultRepoDirname,
	}
}

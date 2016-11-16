package main

import "github.com/urfave/cli"

const (
	OutputPretty = iota
	OutputLog
	OutputUgly
)

const DefaultRepoDirname = ".just"

type Config struct {
	Debug       bool
	Cwd         string
	RepoDirname string
	OutputMode  int
}

func NewConfigWithDefaults() *Config {
	return &Config{
		RepoDirname: DefaultRepoDirname,
	}
}

func NewConfigFromCmdline(c *cli.Context) *Config {
	conf := NewConfigWithDefaults()
	// TODO merge cmdline config with config
	conf.Debug = c.Bool("debug")
	// TODO recurse up to find repo root
	// TODO we should just use the filesystem object to set the mount point
	conf.Cwd = c.String("cwd")
	if c.Bool("pretty") {
		conf.OutputMode = OutputPretty
	} else if c.Bool("plain") {
		conf.OutputMode = OutputLog
	} else if c.Bool("ugly") {
		conf.OutputMode = OutputUgly
	} else {
		conf.OutputMode = getOutputModeFromTTYState()
	}
	return conf
}

func getOutputModeFromTTYState() int {
	// TODO check if we are attached to a tty
	return OutputPretty
}

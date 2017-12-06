package commands

import (
	"github.com/jawher/mow.cli"
	"go-AutoUpdate1C/config"
)

var Commands = []Command{
	&Update{},
	//&List{},
	//&Read{},
	//&Run{},
	//&Task{},
	//&Tasks{},
	//&Top{},
}

type Command interface {
	Name() string
	Desc() string
	Init(config config.Config) func(*cli.Cmd)
}

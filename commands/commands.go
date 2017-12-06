package commands

import (
	"github/Khorevaa/go-AutoUpdate1C/config"

	"github.com/jawher/mow.cli"
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

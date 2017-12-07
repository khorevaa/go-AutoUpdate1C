package commands

import (
	"github.com/khorevaa/go-AutoUpdate1C/config"

	"github.com/jawher/mow.cli"
)

var Commands = []Command{
	&Update{},
	&Sessions{
		subCommands: []Command{
			&SessionsLock{},
			&SessionsUnLock{},
			&SessionsKill{},
		},
	},
	&Run{},
	&Backups{},
	&Agent{},
	//&Tasks{},
	//&Top{},
}

type Command interface {
	Name() string
	Desc() string
	Init(config config.ConfigFn) func(*cli.Cmd)
}

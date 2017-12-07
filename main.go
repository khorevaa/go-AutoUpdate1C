// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/khorevaa/go-AutoUpdate1C/commands"
	"github.com/khorevaa/go-AutoUpdate1C/config"
	"os"
)

func main() {

	app := cli.App("AutoUpdate1C", "Автоматические обновление 1С")

	app.Version("v version", "1.0")

	var (
		debug     = app.BoolOpt("debug", false, "Вывод отладочной информации")
		v8version = app.StringOpt("v8 v8version", "8.3", "Версия платформы 1С.Предприятие")
	)

	var help = app.Bool(cli.BoolOpt{
		Name:      "h help",
		Value:     false,
		Desc:      "Показать справку и выйти",
		HideValue: true,
	})

	app.Before = func() {
		if *help {
			app.PrintLongHelp()
		}

		if *debug {
			fmt.Println("Включен режим отладки")
		}
	}

	app.ErrorHandling = flag.ExitOnError

	config := config.Config{
		V8:    *v8version,
		Debug: *debug,
	}

	for _, cmd := range commands.Commands {
		app.Command(cmd.Name(), cmd.Desc(), cmd.Init(config))
	}
	app.Run(os.Args)

}

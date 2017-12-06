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

//import "go-AutoUpdate1C/cmd"
import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	//"github.com/spf13/cobra"
	"go-AutoUpdate1C/commands"
	"go-AutoUpdate1C/config"
)

func main() {

	app := cli.App("1updater", "Автоматические обновление 1С")
	//app.Spec = "[-v] [--v8]"

	var (
		verbose   = app.BoolOpt("v verbose", false, "Verbose debug mode")
		v8version = app.StringOpt("v8 v8version", "8.3", "Версия платформы 1С.Предприятие")
	)

	var help = app.Bool(cli.BoolOpt{
		Name:      "h help",
		Value:     false,
		Desc:      "Show the help info and exit",
		HideValue: true,
	})

	app.Before = func() {
		if *help {
			app.PrintLongHelp()
		}
		if *verbose {
			// Here you can enable debug output in your logger for example
			fmt.Println("Verbose mode enabled")
		}
	}

	config := config.Config{
		V8: *v8version,
	}

	//app.After = func() {
	//	if workErr != nil {
	//		fmt.Printf("Ошибка выполнения: %v", workErr)
	//		cli.Exit(1)
	//	}
	//
	//}

	for _, cmd := range commands.Commands {
		app.Command(cmd.Name(), cmd.Desc(), cmd.Init(config))
	}
	app.Run(os.Args)

}

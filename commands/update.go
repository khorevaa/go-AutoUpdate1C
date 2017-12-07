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

package commands

import (
	"github.com/khorevaa/go-AutoUpdate1C/config"
	"github.com/khorevaa/go-AutoUpdate1C/update"

	"github.com/khorevaa/go-AutoUpdate1C/logging"

	"github.com/jawher/mow.cli"
)

type Update struct{}

func (_ Update) Name() string { return "update u" }
func (_ Update) Desc() string {
	return "Обновление конфигурации информационной базы"
}

func (_ Update) Init(config config.ConfigFn) func(*cli.Cmd) {

	updateInit := func(cmd *cli.Cmd) {
		// These are the command specific options and args, nicely scoped inside a func
		var (
			loadCf    = cmd.BoolOpt("load-cf l", false, "Выполнить загрузку конфигурации из файла, вместо обновления")
			dbUser    = cmd.StringOpt("db-user w", "Администратор", "Пользователь информационной базы")
			dbPwd     = cmd.StringOpt("db-pwd p", "", "Пароль пользователя информационной базы")
			ucCode    = cmd.StringOpt("uc-code c", "", "Ключ разрешения запуска")
			db        = cmd.StringArg("CONNECT", "", "Строка подключения к информационной базе")
			updateDir = cmd.StringArg("FILE", "", "Путь к файлу обновления (папка или указание на *.cf, *.cfu)")
		)

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			logUpdate := config().Log().NewContextLogger(logging.LogFeilds{
				"command": "update",
			})

			Обновлятор := update.НовоеОбновление(*db, *dbUser, *dbPwd)
			failOnErr(Обновлятор.УстановитьВерсиюПлатформы(config().V8))
			Обновлятор.УстановитьВремяОжидания(config().TimeOut)
			Обновлятор.ФайлОбновления = *updateDir
			Обновлятор.ВыполнитьЗагрузкуВместоОбновения = *loadCf
			Обновлятор.УстановитьЛог(logUpdate)
			Обновлятор.УстановитьКлючРазрешенияЗапуска(*ucCode)
			workErr := Обновлятор.ВыполнитьОбновление()

			if workErr != nil {
				logUpdate.Context(logging.LogFeilds{
					"db":        *db,
					"updateDir": *updateDir,
					"dbUser":    *dbUser,
					"ucCode":    *ucCode,
					"loadCf":    *loadCf,
					"v8":        config().V8,
				}).WithError(workErr).Error("Ошибка выполнения команды")
			}
			failOnErr(workErr)

		}
		//cmd.Spec = "[-l --uc -u -p]"
	}

	return updateInit
}

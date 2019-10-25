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
	"github.com/jawher/mow.cli"
	"github.com/khorevaa/go-AutoUpdate1C/config"
	"github.com/khorevaa/go-AutoUpdate1C/logging"
	"github.com/khorevaa/go-AutoUpdate1C/update"
)

type Backups struct{}

func (_ Backups) Name() string { return "backup b" }
func (_ Backups) Desc() string {
	return "Управление выгрузкой из информационной базы"
}

func (c Backups) Init(config config.ConfigFn) func(*cli.Cmd) {

	sessionsInit := func(cmd *cli.Cmd) {

		var (
			dbUser     = cmd.StringOpt("db-user u", "Администратор", "user информационной базы")
			dbPwd      = cmd.StringOpt("db-pwd p", "", "Password пользователя информационной базы")
			ucCode     = cmd.StringOpt("uc-code c", "", "Ключ разрешения запуска")
			rewrite    = cmd.BoolOpt("rewrite", false, "Перезаписть существующий файл бекапа")
			restore    = cmd.BoolOpt("restore r", false, "Вссстановить информационной базы из выгрузки")
			db         = cmd.StringArg("CONNECT", "", "Строка подключения к информационной базе")
			backUpFile = cmd.StringArg("FILE", "", "Путь к файлу выгрузки из информационной базы")
		)

		cmd.Spec = "[OPTIONS] ( [--restore | --rewrite] ) CONNECT FILE"

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			logCommand := config().Log().NewContextLogger(logging.LogFeilds{
				"command": "backup",
			})
			Обновлятор := update.НовоеОбновление(*db, *dbUser, *dbPwd)
			failOnErr(Обновлятор.УстановитьВерсиюПлатформы(config().V8))
			Обновлятор.УстановитьВремяОжидания(config().TimeOut)
			Обновлятор.УстановитьКлючРазрешенияЗапуска(*ucCode)
			Обновлятор.УстановитьЛог(logCommand)

			var workErr error

			if *restore {
				log := Обновлятор.Log.Context(logging.LogFeilds{
					"ФайлВыгрузки": *backUpFile,
				})
				log.Info("Выполняю загрузку информационной базы")
				workErr = Обновлятор.ЗагрузитьИнформационнуюБазу(*backUpFile)
				log.IfError(workErr, "Ошибка загрузки в информационную базу")

			} else {
				log := Обновлятор.Log.Context(logging.LogFeilds{
					"ФайлВыгрузки": *backUpFile,
					"Перазаписать": *rewrite,
				})
				log.Info("Выполняю выгрузку информационной базы")
				workErr = Обновлятор.ВыгрузитьИнформационнуюБазу(*backUpFile)
				log.IfError(workErr, "Ошибка выгрузки в информационную базу")
			}

			if workErr != nil {
				logCommand.Context(logging.LogFeilds{
					"db":         *db,
					"dbUser":     *dbUser,
					"ucCode":     *ucCode,
					"v8":         config().V8,
					"backUpFile": *backUpFile,
					"rewrite":    *rewrite,
					"restore":    *restore,
				}).WithError(workErr).Error("Ошибка выполнения команды: ")
			}

			failOnErr(workErr)

		}

	}

	return sessionsInit
}

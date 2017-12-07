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

type Run struct{}

func (_ Run) Name() string { return "run r" }
func (_ Run) Desc() string {
	return "Запуск базы даныхх в режиме 1С.Предприятие"
}

func (_ Run) Init(config config.ConfigFn) func(*cli.Cmd) {

	updateInit := func(cmd *cli.Cmd) {
		// These are the command specific options and args, nicely scoped inside a func
		var (
			dbUser     = cmd.StringOpt("db-user u", "Администратор", "Пользователь информационной базы")
			dbPwd      = cmd.StringOpt("db-pwd p", "", "Пароль пользователя информационной базы")
			ucCode     = cmd.StringOpt("uc-code c", "", "Ключ разрешения запуска")
			command    = cmd.StringOpt("command", "", "Команда, при запуске")
			privileged = cmd.BoolOpt("privileged", false, "запуск в режиме привилегированного сеанса")
			db         = cmd.StringArg("CONNECT", "", "Строка подключения к информационной базе")
			fileEpf    = cmd.StringArg("FILE", "", "Путь к файлу обработки/отчета выполняемого при запуске")
		)

		cmd.Spec = "[OPTIONS] CONNECT [FILE]"

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			logCommand := config().Log().NewContextLogger(logging.LogFeilds{
				"command": "run",
			})
			Обновлятор := update.НовоеОбновление(*db, *dbUser, *dbPwd)
			failOnErr(Обновлятор.УстановитьВерсиюПлатформы(config().V8))
			Обновлятор.УстановитьВремяОжидания(config().TimeOut)
			Обновлятор.УстановитьЛог(logCommand)
			Обновлятор.УстановитьКлючРазрешенияЗапуска(*ucCode)
			workErr := Обновлятор.ВыполнитьВРежимеПредприятия(*command, *fileEpf, *privileged)

			if workErr != nil {
				logCommand.Context(logging.LogFeilds{
					"db":         *db,
					"fileEpf":    *fileEpf,
					"dbUser":     *dbUser,
					"ucCode":     *ucCode,
					"command":    *command,
					"privileged": *privileged,
					"v8":         config().V8,
				}).WithError(workErr).Error("Ошибка выполнения команды: ")
			}
			failOnErr(workErr)
		}
		//cmd.Spec = "[-l --uc -u -p]"
	}

	return updateInit
}

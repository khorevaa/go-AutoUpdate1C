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
	"github.com/pkg/errors"
	"time"
)

type Agent struct{}

func (_ Agent) Name() string { return "agent a" }
func (_ Agent) Desc() string { return "Запуск в режиме агента обновления" }

func (c Agent) Init(config config.Config) func(*cli.Cmd) {

	commandInit := func(cmd *cli.Cmd) {

		cmd.LongDesc = `Данный режим работает по HTTP (REST API) с базой данных.
		Возможности:
		* самостоятельно получает список информационных баз к обновления;
		* поддержание нескольких потоков обновления
		* переодический/разовый опрос необходимости обновления
		* отправка журнала обновления на url.`

		var (
			restUser = cmd.StringOpt("rest-user u", "", "Пользователь для подключения к серверу REST API")
			restPwd  = cmd.StringOpt("rest-pwd p", "", "Пароль пользователя для подключения к серверу REST API")
			threads  = cmd.IntOpt("processes c", 1, "Количество одновременно работающих процесссов")
			server   = cmd.StringArg("SERVER", "", "Сервер с REST API для получения списка и настроек обновления информационных баз")
		)

		duration := Duration(60)
		cmd.VarOpt("timer t", &duration, "Переодичность опроса сервера REST API в минутах (0 - отключено)")

		logCommand := config.Log().NewContextLogger(logging.LogFeilds{
			"command": "agent",
		})

		//cmd.Spec ="[-u -p -c] ( [-r | --rewrite] ) CONNECT FILE"

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			workErr := errors.New("Команда не реализована") //Обновлятор.ВыполнитьОбновление()

			if workErr != nil {
				logCommand.Context(logging.LogFeilds{
					"server":   *server,
					"restUser": *restUser,
					"restPwd":  *restPwd,
					"threads":  *threads,
					"duration": duration,
				}).WithError(workErr).Error("Ошибка выполнения команды: ")
			}

			failOnErr(workErr)

		}

	}

	return commandInit
}

// Declare your type
type Duration time.Duration

// Make it implement flag.Value
func (d *Duration) Set(v string) error {
	parsed, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*d = Duration(parsed)
	return nil
}

func (d *Duration) String() string {
	duration := time.Duration(*d)
	return duration.String()
}

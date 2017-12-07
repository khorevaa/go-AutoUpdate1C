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

	"time"

	"github.com/khorevaa/go-AutoUpdate1C/commands/types"
	"github.com/pkg/errors"
)

type Sessions struct {
	subCommands []Command
}

type SessionsLock struct {
	subCommands []Command
}
type SessionsUnLock struct {
	subCommands []Command
}

type SessionsKill struct {
	subCommands []Command
}

func (_ Sessions) Name() string { return "sessions s" }
func (_ Sessions) Desc() string {
	return "Управление сеансами в информационной базы"
}

func (c Sessions) Init(config config.ConfigFn) func(*cli.Cmd) {

	sessionsInit := func(cmd *cli.Cmd) {
		// These are the command specific options and args, nicely scoped inside a func

		for _, subCommand := range c.subCommands {
			cmd.Command(subCommand.Name(), subCommand.Desc(), subCommand.Init(config))
		}

		//cmd.Spec = "[-l --uc -u -p]"
	}

	return sessionsInit
}

func (_ SessionsUnLock) Name() string { return "unlock ul" }
func (_ SessionsUnLock) Desc() string { return "Снять блокировку соединений" }
func (c SessionsUnLock) Init(config config.ConfigFn) func(*cli.Cmd) {

	sessionsInit := func(cmd *cli.Cmd) {
		// These are the command specific options and args, nicely scoped inside a func

		for _, subCommand := range c.subCommands {
			cmd.Command(subCommand.Name(), subCommand.Desc(), subCommand.Init(config))
		}

		var (
			dbUser       = cmd.StringOpt("db-user u", "Администратор", "Пользователь информационной базы")
			dbPwd        = cmd.StringOpt("db-pwd p", "", "Пароль пользователя информационной базы")
			ucCode       = cmd.StringOpt("uc-code c", "", "Ключ разрешения запуска")
			clusterAdmin = cmd.StringOpt("cluster-admin", "", "Администратор кластера")
			clusterPwd   = cmd.StringOpt("cluster-pwd", "", "Пароль администратора кластера")
			rasRunMode   = cmd.StringOpt("ras-run r", "noRun", "Режим запуска RAS (noRun, run, stop")
			db           = cmd.StringArg("CONNECT", "", "Строка подключения к информационной базе")
			ras          = cmd.StringArg("RAS", "localhost:1545", "Сетевой адрес RAS")
		)
		cmd.Spec = "[OPTIONS] CONNECT [RAS]"

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			logCommand := config().Log().NewContextLogger(logging.LogFeilds{
				"command":    "sessions",
				"subcommand": "unlock",
			})
			Обновлятор := update.НовоеОбновление(*db, *dbUser, *dbPwd)
			failOnErr(Обновлятор.УстановитьВерсиюПлатформы(config().V8))
			Обновлятор.УстановитьВремяОжидания(config().TimeOut)
			Обновлятор.УстановитьЛог(logCommand)
			Обновлятор.УстановитьКлючРазрешенияЗапуска(*ucCode)
			workErr := errors.New("Команда не реализована") // Обновлятор.ВыполнитьОбновление()

			if workErr != nil {
				logCommand.Context(logging.LogFeilds{
					"db":           *db,
					"dbUser":       *dbUser,
					"ucCode":       *ucCode,
					"v8":           config().V8,
					"ras":          *ras,
					"rasRunMode":   *rasRunMode,
					"clusterAdmin": *clusterAdmin,
					"clusterPwd":   *clusterPwd,
				}).WithError(workErr).Error("Ошибка выполнения команды")
			}
			failOnErr(workErr)

		}
		//cmd.Spec = "[-l --uc -u -p]"
	}

	return sessionsInit
}

func (_ SessionsLock) Name() string { return "lock l" }
func (_ SessionsLock) Desc() string {
	return "Установить блокировку соединений"
}
func (c SessionsLock) Init(config config.ConfigFn) func(*cli.Cmd) {

	return func(cmd *cli.Cmd) {
		// These are the command specific options and args, nicely scoped inside a func

		for _, subCommand := range c.subCommands {
			cmd.Command(subCommand.Name(), subCommand.Desc(), subCommand.Init(config))
		}

		lockStart := types.DateTime{time.Now().Add(5 * time.Minute)}
		lockEnd := types.DateTime{lockStart.Add(1 * time.Hour)}
		var (
			dbUser       = cmd.StringOpt("db-user u", "Администратор", "Пользователь информационной базы")
			dbPwd        = cmd.StringOpt("db-pwd p", "", "Пароль пользователя информационной базы")
			ucCode       = cmd.StringOpt("uc-code c", "", "Ключ разрешения запуска")
			clusterAdmin = cmd.StringOpt("cluster-admin", "", "Администратор кластера")
			clusterPwd   = cmd.StringOpt("cluster-pwd", "", "Пароль администратора кластера")
			lockMessage  = cmd.StringOpt("lock-message m", "Выполняются технические работы. Установлена блокировка сеансов.", "Ключ разрешения запуска")
			rasRunMode   = cmd.StringOpt("ras-run r", "noRun", "Режим запуска RAS (noRun, run, stop")
			db           = cmd.StringArg("CONNECT", "", "Строка подключения к информационной базе")
			ras          = cmd.StringArg("RAS", "localhost:1545", "Сетевой адрес RAS")
		)
		cmd.Spec = "[OPTIONS] CONNECT [RAS]"
		cmd.VarOpt("lock-start s", &lockStart, "Время старта блокировки пользователей, время указываем как '2040-12-31T23:59:59")
		cmd.VarOpt("lock-end e", &lockEnd, "Время окончания блокировки пользователей, время указываем как '2040-12-31T23:59:59")

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			logCommand := config().Log().NewContextLogger(logging.LogFeilds{
				"command":    "sessions",
				"subcommand": "lock",
			})
			Обновлятор := update.НовоеОбновление(*db, *dbUser, *dbPwd)
			failOnErr(Обновлятор.УстановитьВерсиюПлатформы(config().V8))
			Обновлятор.УстановитьВремяОжидания(config().TimeOut)
			Обновлятор.УстановитьЛог(logCommand)
			workErr := errors.New("Команда не реализована") //Обновлятор.ВыполнитьОбновление()

			if workErr != nil {
				logCommand.Context(logging.LogFeilds{
					"db":           *db,
					"dbUser":       *dbUser,
					"ucCode":       *ucCode,
					"v8":           config().V8,
					"ras":          *ras,
					"rasRunMode":   *rasRunMode,
					"lockMessage":  *lockMessage,
					"lockStart":    lockStart.Time.String(),
					"lockEnd":      lockEnd.Time.String(),
					"clusterAdmin": *clusterAdmin,
					"clusterPwd":   *clusterPwd,
				}).WithError(workErr).Error("Ошибка выполнения команды")
			}
			failOnErr(workErr)
		}

	}

}

func (_ SessionsKill) Name() string { return "kill k" }
func (_ SessionsKill) Desc() string {
	return "Удалить все текущие соединения"
}
func (_ SessionsKill) Init(config config.ConfigFn) func(*cli.Cmd) {

	sessionsInit := func(cmd *cli.Cmd) {
		// These are the command specific options and args, nicely scoped inside a func

		var (
			dbUser       = cmd.StringOpt("db-user u", "Администратор", "Пользователь информационной базы")
			dbPwd        = cmd.StringOpt("db-pwd p", "", "Пароль пользователя информационной базы")
			ucCode       = cmd.StringOpt("uc-code c", "", "Ключ разрешения запуска")
			clusterAdmin = cmd.StringOpt("cluster-admin", "", "Администратор кластера")
			clusterPwd   = cmd.StringOpt("cluster-pwd", "", "Пароль администратора кластера")
			rasRunMode   = cmd.StringOpt("ras-run r", "noRun", "Режим запуска RAS (noRun, run, stop")
			db           = cmd.StringArg("CONNECT", "", "Строка подключения к информационной базе")
			ras          = cmd.StringArg("RAS", "localhost:1545", "Сетевой адрес RAS")
		)
		cmd.Spec = "[OPTIONS] CONNECT [RAS]"

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			logCommand := config().Log().NewContextLogger(logging.LogFeilds{
				"command":    "sessions",
				"subcommand": "kill",
			})
			Обновлятор := update.НовоеОбновление(*db, *dbUser, *dbPwd)
			failOnErr(Обновлятор.УстановитьВерсиюПлатформы(config().V8))
			Обновлятор.УстановитьВремяОжидания(config().TimeOut)
			Обновлятор.УстановитьЛог(logCommand)
			workErr := errors.New("Команда не реализована") //Обновлятор.ВыполнитьОбновление()

			if workErr != nil {
				logCommand.Context(logging.LogFeilds{
					"db":           *db,
					"dbUser":       *dbUser,
					"ucCode":       *ucCode,
					"v8":           config().V8,
					"ras":          *ras,
					"rasRunMode":   *rasRunMode,
					"clusterAdmin": *clusterAdmin,
					"clusterPwd":   *clusterPwd,
				}).WithError(workErr).Error("Ошибка выполнения команды")
			}
			failOnErr(workErr)

		}
		//cmd.Spec = "[-l --uc -u -p]"
	}

	return sessionsInit
}

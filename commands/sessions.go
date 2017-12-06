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
	"github/Khorevaa/go-AutoUpdate1C/config"
	"github/Khorevaa/go-AutoUpdate1C/update"

	"github/Khorevaa/go-AutoUpdate1C/logging"

	"github.com/jawher/mow.cli"

	"time"
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

func (_ Sessions) Name() string { return "sessions s" }
func (_ Sessions) Desc() string {
	return "Управление сеансами в информационной базы"
}

func (c Sessions) Init(config config.Config) func(*cli.Cmd) {

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
func (_ SessionsUnLock) Desc() string {
	return "Снять блокировку соединений"
}

func (c SessionsUnLock) Init(config config.Config) func(*cli.Cmd) {

	return func(cmd *cli.Cmd) {
		// These are the command specific options and args, nicely scoped inside a func

		for _, subCommand := range c.subCommands {
			cmd.Command(subCommand.Name(), subCommand.Desc(), subCommand.Init(config))
		}

		lockStart := Duration(0)
		var (
			dbUser    = cmd.StringOpt("db-user w", "Администратор", "Пользователь информационной базы")
			dbPwd     = cmd.StringOpt("db-pwd p", "", "Пароль пользователя информационной базы")
			ucCode    = cmd.StringOpt("uc-code uc", "", "Ключ разрешения запуска")
			rac       = cmd.StringOpt("rac r", "localhost:1545", "Сетевой адрес RAS")
			updateDir = cmd.StringArg("CONNECT", "", "Строка подключения к информационной базе")

			lockMessage = cmd.StringOpt("lock-message m", "", "Ключ разрешения запуска")
		)

		cmd.VarOpt("lock-start s", &lockStart, "Время старта блокировки пользователей, время указываем как '2040-12-31T23:59:59")
		cmd.VarOpt("lock-end s", &lockStart, "Время старта блокировки пользователей, время указываем как '2040-12-31T23:59:59")

		logUpdate := config.Log().NewContextLogger(logging.LogFeilds{
			"command": "update",
		})

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			Обновлятор := update.НовоеОбновление(*db, *dbUser, *dbPwd, *ucCode)
			Обновлятор.УстановитьВерсиюПлатформы(config.V8)
			Обновлятор.ФайлОбновления = *updateDir
			Обновлятор.ВыполнитьЗагрузкуВместоОбновения = *loadCf
			Обновлятор.УстановитьЛог(logUpdate)
			workErr := Обновлятор.ВыполнитьОбновление()

			if workErr != nil {
				logUpdate(logging.LogFeilds{
					"db":        *db,
					"updateDir": *updateDir,
					"loadCf":    *loadCf,
					"ucCode":    *ucCode,
					"v8":        config.V8,
				}).WithError(workErr).Error("Ошибка выполнения команды: ")
			}

		}
		//cmd.Spec = "[-l --uc -u -p]"
	}

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

func (_ SessionsLock) Name() string { return "lock l" }
func (_ SessionsLock) Desc() string {
	return "Установить блокировку соединений"
}

func (c SessionsLock) Init(config config.Config) func(*cli.Cmd) {

	sessionsInit := func(cmd *cli.Cmd) {
		// These are the command specific options and args, nicely scoped inside a func

		for _, subCommand := range c.subCommands {
			cmd.Command(subCommand.Name(), subCommand.Desc(), subCommand.Init(config))
		}

		var (
			loadCf    = cmd.BoolOpt("load-cf l", false, "Выполнить загрузку конфигурации из файла, вместо обновления")
			dbUser    = cmd.StringOpt("db-user w", "Администратор", "Пользователь информационной базы")
			dbPwd     = cmd.StringOpt("db-pwd p", "", "Пароль пользователя информационной базы")
			ucCode    = cmd.StringOpt("uc-code uc", "", "Ключ разрешения запуска")
			db        = cmd.StringArg("CONNECT", "", "Строка подключения к информационной базе")
			updateDir = cmd.StringArg("FILE", "", "Путь к файлу обновления (папка или указание на *.cf, *.cfu)")
		)

		logUpdate := config.Log().NewContextLogger(logging.LogFeilds{
			"command": "update",
		})

		// What to run when this command is called
		cmd.Action = func() {
			// Inside the action, and only inside, you can safely access the values of the options and arguments

			Обновлятор := update.НовоеОбновление(*db, *dbUser, *dbPwd, *ucCode)
			Обновлятор.УстановитьВерсиюПлатформы(config.V8)
			Обновлятор.ФайлОбновления = *updateDir
			Обновлятор.ВыполнитьЗагрузкуВместоОбновения = *loadCf
			Обновлятор.УстановитьЛог(logUpdate)
			workErr := Обновлятор.ВыполнитьОбновление()

			if workErr != nil {
				logUpdate(logging.LogFeilds{
					"db":        *db,
					"updateDir": *updateDir,
					"loadCf":    *loadCf,
					"ucCode":    *ucCode,
					"v8":        config.V8,
				}).WithError(workErr).Error("Ошибка выполнения команды: ")
			}

		}
		//cmd.Spec = "[-l --uc -u -p]"
	}

	return sessionsInit
}

func init() {
	RootCmd.AddCommand(sessionsCmd)

	sessionsCmd.AddCommand(lockDBCmd)
	sessionsCmd.AddCommand(unlockDBCmd)
	// Here you will define your flags and configuration settings.
	sessionsCmd.PersistentFlags().StringP("ras", "r", "localhost:1545", "Сетевой адрес RAS")

	sessionsCmd.PersistentFlags().StringP("db", "c", "", "Строка подключения к информационной базе")
	sessionsCmd.PersistentFlags().StringP("db-user", "u", "", "Пользователь информационной базы")
	sessionsCmd.PersistentFlags().StringP("db-pwd", "p", "", "Пароль пользователя информационной базы")

	sessionsCmd.PersistentFlags().StringP("cluster-admin", "", "", "Администратор кластера")
	sessionsCmd.PersistentFlags().StringP("cluster-pwd", "", "", "Пароль администратора кластера")

	lockDBCmd.Flags().StringP("lock-message", "", "", "Сообщение облокировки")
	lockDBCmd.Flags().StringP("lock-uc-code", "", "", "Ключ разрешения запуска")
	lockDBCmd.Flags().StringP("lock-start", "", "", "Время старта блокировки пользователей, время указываем как '2040-12-31T23:59:59'")
	lockDBCmd.Flags().StringP("lock-started", "", "", "Время старта блокировки через n сек")

	unlockDBCmd.Flags().StringP("uc-code", "", "", "Ключ разрешения запуска")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sessionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sessionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

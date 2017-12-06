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

package cmd

import (
	"github.com/Khorevaa/go-AutoUpdate1C/update"
	"github.com/spf13/cobra"
)

var db, dbUser, dbPwd, ucCode string
var updateFile string
var loadCf bool

// updateDBCmd represents the updateDB command
var updateDBCmd = &cobra.Command{
	Use:   "update",
	Short: "Обновление конфигурации информационной базы",
	Long: `Команда производит обновление конфигурацию информационной базы
	Возможно загрука конфигурации в базу данных, вместо обновления`,
	RunE: updateFunc,
}

func updateFunc(cmd *cobra.Command, args []string) (err error) {

	Обновлятор := update.НовоеОбновление(db, dbUser, dbPwd, ucCode)
	Обновлятор.УстановитьВерсиюПлатформы(v8version)
	Обновлятор.ФайлОбновления = updateFile
	Обновлятор.ВыполнитьЗагрузкуВместоОбновения = loadCf

	err = Обновлятор.ВыполнитьОбновление()

	return
}

func init() {
	RootCmd.AddCommand(updateDBCmd)

	updateDBCmd.Flags().StringVarP(&db, "db", "c", "", "Строка подключения к информационной базе")
	updateDBCmd.Flags().StringVarP(&dbUser, "db-user", "u", "", "Пользователь информационной базы")
	updateDBCmd.Flags().StringVarP(&dbPwd, "db-pwd", "p", "", "Пароль пользователя информационной базы")

	updateDBCmd.Flags().StringVarP(&updateFile, "update-file", "f", "", "Путь к файлу обновления (папка или указание на *.cf, *.cfu)")
	updateDBCmd.Flags().StringVarP(&ucCode, "uc-code", "", "", "Ключ разрешения запуска")

	updateDBCmd.Flags().BoolVarP(&loadCf, "load-cf", "l", false, "Выполнить загрузку конфигурации из файла, вместо обновления")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateDBCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateDBCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

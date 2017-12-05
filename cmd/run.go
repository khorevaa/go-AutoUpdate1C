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
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Запуск базы даныхх в режиме 1С.Предприятие",
	Long: `Запуск в релиме 1С.Предприятие позволяет выполнипть
	действия после обновления в пользовательском режиме.
	Так же возможен запуск обработок`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

func init() {
	RootCmd.AddCommand(runCmd)


	runCmd.Flags().StringP("db", "c", "", "Строка подключения к информационной базе")
	runCmd.Flags().StringP("db-user", "u", "", "Пользователь информационной базы")
	runCmd.Flags().StringP("db-pwd", "p", "", "Пароль пользователя информационной базы")

	runCmd.Flags().StringP("uc-code", "", "", "Ключ разрешения запуска")

	runCmd.Flags().StringP("epf-file", "e", "", "Путь к файлу обработки (*.epf)")


}

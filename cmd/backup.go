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

// backupDBCmd represents the backupDB command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Создание бекап информационной базы данных",
	Long: `Команда создает бекап информационной базы данных и размещает его в указанном месте`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backupDB called")
	},
}

func init() {

	RootCmd.AddCommand(backupCmd)

	backupCmd.PersistentFlags().StringP("db", "c", "", "Строка подключения к информационной базе")
	backupCmd.PersistentFlags().StringP("db-user", "u", "", "Пользователь информационной базы")
	backupCmd.PersistentFlags().StringP("db-pwd", "p", "", "Пароль пользователя информационной базы")

	backupCmd.Flags().StringP("backup-file", "f", "", "Путь к файлу бекапа")
	backupCmd.Flags().BoolP("use-temp-file", "", false, "Делать промежуточный файл букапа, а потом копировать в укзанное место")
	backupCmd.Flags().BoolP("rewrite-file", "", true, "Флаг перезаписи существующего файла бекапа")

	backupCmd.Flags().StringP("uc-code", "", "", "Ключ разрешения запуска")

}

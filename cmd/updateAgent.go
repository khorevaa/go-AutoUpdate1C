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

// updateAgentCmd represents the updateAgent command
var updateAgentCmd = &cobra.Command{
	Use:   "updateAgent",
	Short: "Режим запуска агента обновления",
	Long: `Данный режим работает по HTTP (REST API) с базой данных.
Возможности:
	* самостоятельно получает список информационных баз к обновления;
	* поддержание нескольких потоков обновления
	* переодический/разовый опрос необходимости обновления
	* отправка журнала обновления на url.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateAgent called")
	},
}

func init() {
	RootCmd.AddCommand(updateAgentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateAgentCmd.PersistentFlags().String("foo", "", "A help for foo")

	updateAgentCmd.Flags().StringP("rest-url", "r", "", "Сервер с REST API для получения списка и настроек обновления информационных баз")
	updateAgentCmd.Flags().StringP("rest-user", "u", "", "Пользователь для подключения к серверу REST API")
	updateAgentCmd.Flags().StringP("rest-pwd", "p", "", "Пароль пользователя для подключения к серверу REST API")
	updateAgentCmd.Flags().Int8P("timer", "t", 0, "Переодичность опроса сервера REST API в минутах (0 - отключено)")

	updateAgentCmd.Flags().Int8P("count-threads", "d", 1, "Количество одновременно работающих процесссов (0 - отключено)")

}

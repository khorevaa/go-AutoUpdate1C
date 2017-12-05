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

// updateDBCmd represents the updateDB command
var updateDBCmd = &cobra.Command{
	Use:   "updateDB",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateDB called")
	},
}

func init() {
	RootCmd.AddCommand(updateDBCmd)

	var СтрокаСоединенияСБазой string
	updateDBCmd.Flags().StringVarP(&СтрокаСоединенияСБазой,"db-path", "c", "", "строка подключения к базе данных")
	updateDBCmd.Flags().StringVarP(&СтрокаСоединенияСБазой,"user", "u", "", "пользователь для подключения к базе данных")
	updateDBCmd.Flags().StringVarP(&СтрокаСоединенияСБазой,"password", "p", "", "пароль пользователя для подключения к базе данных")

	updateDBCmd.Flags().StringVarP(&СтрокаСоединенияСБазой,"update-file", "f", "", "Путь к файлу обновления (папка или указание на *.cf, *.cfu)")
	updateDBCmd.Flags().StringVarP(&СтрокаСоединенияСБазой,"uc-code", "", "", "код доступа к заблокированной базе данных (/UC)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateDBCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateDBCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

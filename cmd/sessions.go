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

// sessionsCmd represents the sessions command
var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "Управление сеансами в базе данных",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sessions called")
	},
}

// lockDBCmd represents the lockDB command
var lockDBCmd = &cobra.Command{
	Use:   "lock",
	Short: "Установить блокировку соединений с базой данных",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lockDB called")
	},
}

// unlockDBCmd represents the unlockDB command
var unlockDBCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Снять блокировку соединений с базой данных",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unlockDB called")
	},
}

func init() {
	RootCmd.AddCommand(sessionsCmd)

	sessionsCmd.AddCommand(lockDBCmd)
	sessionsCmd.AddCommand(unlockDBCmd)
	// Here you will define your flags and configuration settings.

	sessionsCmd.PersistentFlags().StringP("db", "c", "", "A help for foo")
	sessionsCmd.PersistentFlags().StringP("db-user", "u", "", "Пользователь информационной базы")
	sessionsCmd.PersistentFlags().StringP("db-pwd", "p", "", "Пароль пользователя информационной базы")

	sessionsCmd.PersistentFlags().StringP("cluster-admin", "", "", "Администратор кластера")
	sessionsCmd.PersistentFlags().StringP("cluster-pwd", "", "", "Пароль администратора кластера")

	lockDBCmd.Flags().StringP("lock-message", "", "", "Сообщение облокировки")
	lockDBCmd.Flags().StringP("lock-uc-code", "", "", "Ключ разрешения запуска")
	lockDBCmd.Flags().StringP("lock-start", "", "", "Время старта блокировки пользователей, время указываем как '2040-12-31T23:59:59'")
	lockDBCmd.Flags().StringP("lock-started", "", "", "Время старта блокировки через n сек")

	sessionsCmd.PersistentFlags().StringP("cluster-pwd", "", "", "A help for foo")

	sessionsCmd.Flags().String("command", "", "Команда действия")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sessionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sessionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

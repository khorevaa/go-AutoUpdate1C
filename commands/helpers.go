package commands

import (
	"bytes"
	"fmt"

	"github.com/jawher/mow.cli"
)

func truncStr(s string, l int) string {
	runes := bytes.Runes([]byte(s))
	if len(runes) < l {
		return s
	}
	return string(runes[:l])
}

func failOnErr(err error) {
	if err != nil {
		fmt.Printf("Ошибка выполненния программы: %v \n", err.Error())
		cli.Exit(1)
	}
}

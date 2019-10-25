package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	worker "github.com/khorevaa/go-AutoUpdate1C/worker/cmd"
	"github.com/vito/twentythousandtonnesofcrudeoil"
	"os"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

type WebCommand struct {
}

type AutoUpdateCommand struct {
	Version func() `short:"v" long:"version" description:"Print the version of AutoUpdate1C and exit"`

	Web    WebCommand           `command:"web"     description:"Run the web UI and work scheduler."`
	Worker worker.WorkerCommand `command:"worker"  description:"Run and register a worker."`
}

func main() {
	var cmd AutoUpdateCommand

	cmd.Version = func() {
		fmt.Println(buildVersion(version, commit, date, builtBy))
		os.Exit(0)
	}

	parser := flags.NewParser(&cmd, flags.Default)
	parser.NamespaceDelimiter = "-"

	twentythousandtonnesofcrudeoil.TheEnvironmentIsPerfectlySafe(parser, "AU1C_")

	_, err := parser.Parse()
	if err != nil {
		//fmt.Fprintln(os.Stderr, err)
		parser.WriteHelp(os.Stderr)
		os.Exit(1)

	}
}

func buildVersion(version, commit, date, builtBy string) string {
	var result = fmt.Sprintf("version: %s", version)
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}

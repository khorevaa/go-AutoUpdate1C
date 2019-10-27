package v8run

import (
	"os/exec"
	"strconv"
	"syscall"

	"context"
	"github.com/khorevaa/go-v8runner/v8tools"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	COMMANE_DESIGNER             = "DESIGNER"
	COMMAND_CREATEINFOBASE       = "CREATEINFOBASE"
	COMMAND_ENTERPRISE           = "ENTERPRISE"
	DEFAULT_1SSERVER_PORT  int16 = 1541
)

const (
	EXITCODE_SUCCESS        = 0   // команда выполнена успешно.
	EXITCODE_FAIL           = 1   // при выполнении команды произошла ошибка.
	EXITCODE_DATABASE_ERROR = 101 // при выполнении команды обнаружены ошибки в данных.

)

var ERROR_CHECK_RUNNING = errors.New("error checking runeble")
var ERROR_VERSION_NOT_FOUND = errors.New("error Version is not found")
var ERROR_RUNNING_TIMEOUT = errors.New("error running Timeout executed")

var ERROR_RUNNING_FAILED = errors.New("error running v8 fail")
var ERROR_RUNNING_DATABASE_ERROR = errors.New("error running v8 database error")

type Option func(options *RunOptions)

func WithTimeout(timeout time.Duration) Option {
	return func(r *RunOptions) {
		r.Timeout = timeout
	}
}

func WithOut(file string) Option {
	return func(r *RunOptions) {
		r.Out = file
	}
}

func WithDumpResult(file string) Option {
	return func(r *RunOptions) {
		r.DumpResult = file
	}
}

func WithVersion(version string) Option {
	return func(r *RunOptions) {
		r.Version = version
	}
}

type RunOptions struct {
	Version    string
	Timeout    time.Duration
	Out        string
	DumpResult string
	v8path     string
}

func getV8Path(options *RunOptions) (string, error) {
	if len(options.v8path) > 0 {
		return options.v8path, nil
	}

	return "", ERROR_VERSION_NOT_FOUND
}

func readOut(file string) (string, error) {
	b, err := readV8file(file)
	return string(b), err
}

func readDumpResult(file string) int {

	b, _ := readV8file(file)
	code, _ := strconv.ParseInt(string(b), 10, 64)
	return int(code)
}

func RunWithOptions(where Where, what Command, options *RunOptions) (int, error) {

	var exitCode int

	if what.Check() {
		return EXITCODE_FAIL, ERROR_CHECK_RUNNING
	}

	commandV8, err := getV8Path(options)

	if err != nil {
		return EXITCODE_FAIL, err
	}

	// Create a new context and add a Timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	var args []string
	args = append(args, what.Command())
	args = append(args, where.ConnectString())
	args = append(append(args, what.Args()...))

	cmd := exec.CommandContext(ctx, commandV8, args...)

	out, runErr := cmd.Output()

	if runErr != nil {
		// try to get the exit code
		if exitError, ok := runErr.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			log.Debugf("Could not get exit code for failed program: %v, %v", commandV8, args)
			exitCode = EXITCODE_FAIL
		}
	} else {

		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()

	}

	// We want to Check the context error to see if the Timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		return EXITCODE_FAIL, ERROR_RUNNING_TIMEOUT
	}

	dumpCode := readDumpResult(options.DumpResult)

	switch dumpCode {

	case EXITCODE_SUCCESS:

		return EXITCODE_SUCCESS, nil

	case EXITCODE_FAIL:

		return EXITCODE_FAIL, ERROR_RUNNING_FAILED

	case EXITCODE_DATABASE_ERROR:

		return EXITCODE_DATABASE_ERROR, ERROR_RUNNING_DATABASE_ERROR

	default:

		return EXITCODE_FAIL, errors.New("unknown error")
	}
}

func defaultOptions(command string) *RunOptions {

	var options RunOptions
	options.Timeout = time.Duration(120)

	switch command {

	case COMMANE_DESIGNER:
	case COMMAND_ENTERPRISE:
	case COMMAND_CREATEINFOBASE:
	}

	return &options
}

type Where interface {
	ConnectString() string
}

type Command interface {
	Command() string
	Args() []string
	Check() bool
}

func Run(where Where, what Command, opts ...Option) (int, error) {

	options := defaultOptions(what.Command())

	for _, opt := range opts {
		opt(options)
	}

	return RunWithOptions(where, what, options)

}

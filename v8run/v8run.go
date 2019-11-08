package v8run

import (
	"fmt"
	"github.com/khorevaa/go-AutoUpdate1C/v8run/types"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"context"
	"github.com/pkg/errors"
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

var VERSION_1S = "8.3"

type Option func(options *RunOptions)

func WithTimeout(timeout int64) Option {
	return func(r *RunOptions) {
		r.Timeout = timeout

		if r.Context == nil {
			r.Context = context.Background()
		}

	}
}

func WithContext(ctx context.Context) Option {
	return func(r *RunOptions) {
		r.Context = ctx
	}
}

func WithOut(file string, noTruncate bool) Option {
	return func(r *RunOptions) {
		r.Out = file
		r.tempOut = false
		r.NoTruncate = noTruncate
	}
}

func WithPath(path string) Option {
	return func(r *RunOptions) {
		r.v8path = path
	}
}

func WithDumpResult(file string) Option {
	return func(r *RunOptions) {
		r.DumpResult = file
		r.tempDumpResult = false
	}
}

func WithVersion(version string) Option {
	return func(r *RunOptions) {
		r.Version = version
	}
}

type RunOptions struct {
	Version        string
	Timeout        int64
	Out            string
	NoTruncate     bool
	tempOut        bool
	DumpResult     string
	tempDumpResult bool
	v8path         string
	Context        context.Context
}

func (ro *RunOptions) NewOutFile() {

	tempLog, _ := ioutil.TempFile("", "v8_log_*.txt")

	ro.Out = tempLog.Name()
	ro.tempOut = true

}

func (ro *RunOptions) RemoveOutFile() {

	_ = os.Remove(ro.Out)

}

func (ro *RunOptions) NewDumpResultFile() {

	tempLog, _ := ioutil.TempFile("", "v8_DumpResult_*.txt")

	ro.DumpResult = tempLog.Name()
	ro.tempDumpResult = true

}

func (ro *RunOptions) RemoveDumpResultFile() {

	_ = os.Remove(ro.DumpResult)

}

func (ro *RunOptions) RemoveTempFiles() {

	if ro.tempDumpResult {
		_ = os.Remove(ro.DumpResult)
	}

	if ro.tempOut {
		_ = os.Remove(ro.Out)
	}

}

func getV8Path(options *RunOptions) (string, error) {
	if len(options.v8path) > 0 {
		return options.v8path, nil
	}

	v8 := VERSION_1S
	if len(options.Version) > 0 {
		v8 = options.Version
	}

	fmt.Println(v8)

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

func runCommand(command string, args []string, opts *RunOptions) (err error) {

	cmd := exec.Command(command, args...)
	err = cmd.Run()

	return
}

func runCommandContext(ctx context.Context, command string, args []string, opts *RunOptions) (err error) {

	// Create a new context and add a Timeout to it

	runCtx := ctx
	if opts.Timeout > 0 {

		timeout := int64(time.Second) * opts.Timeout
		ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout))
		defer cancel() // The cancel should be deferred so resources are cleaned up
		runCtx = ctx
	}
	cmd := exec.CommandContext(runCtx, command, args...)

	err = cmd.Run()

	return
}

func RunWithOptions(where types.InfoBase, what types.Command, options *RunOptions) (int, error) {

	if !what.Check() {
		return EXITCODE_FAIL, ERROR_CHECK_RUNNING
	}

	commandV8, err := getV8Path(options)

	if err != nil {
		return EXITCODE_FAIL, err
	}

	var args []string
	args = append(args, what.Command())

	connectString := where.ShortConnectString()
	if what.Command() == COMMAND_CREATEINFOBASE {
		connectString, err = where.CreateString()
	}

	if err != nil {
		return EXITCODE_FAIL, err
	}

	args = append(args, connectString)
	args = append(args, processUserOptions(what.Values())...)

	args = append(args, fmt.Sprintf("/Out %s", options.Out))

	if options.NoTruncate {
		args = append(args, "-NoTruncate")
	}

	args = append(args, fmt.Sprintf("/DumpResult %s", options.DumpResult))
	defer options.RemoveTempFiles()

	var errRun error
	if options.Context != nil {
		errRun = runCommandContext(options.Context, commandV8, args, options)
	} else {
		errRun = runCommand(commandV8, args, options)
	}

	if options.Context != nil && options.Context.Err() == context.DeadlineExceeded {
		return EXITCODE_FAIL, ERROR_RUNNING_TIMEOUT
	}

	if errRun != nil {
		return EXITCODE_FAIL, ERROR_RUNNING_FAILED
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

	options := RunOptions{}

	options.NewOutFile()
	options.NewDumpResultFile()

	return &options
}

func Run(where types.InfoBase, what types.Command, opts ...Option) (int, error) {

	options := defaultOptions(what.Command())

	for _, opt := range opts {
		opt(options)
	}

	return RunWithOptions(where, what, options)

}

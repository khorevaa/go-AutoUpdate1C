package v8run

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"strconv"
	"syscall"

	"context"
	"github.com/khorevaa/go-v8runner/v8platform"
	"github.com/khorevaa/go-v8runner/v8tools"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type ЗапускательИнтерфейс interface {
	ВыполнитьКомандуКонфигуратора() (err error)
	ВыполнитьКомандуСоздатьБазу() (err error)
	ВыполнитьКомандуПредприятие() (err error)
	ВыполнитьКоманду() (err error)

	УстановитьВерсиюПлатформы(строкаВерсияПлатформы string)
	КлючСоединенияСБазой() string
	УстановитьКлючСоединенияСБазой(КлючСоединенияСБазой string)
	УстановитьАвторизацию(Пользователь string, Пароль string)
	УстановитьПараметры(Параметры ...string)
	ДобавитьПараметры(Параметры ...string)
	ПолучитьВыводКоманды() (s string)
	ПроверитьВозможностьВыполнения() (err error)
}

type runeble interface {
	command() string
	args() []string
	check() bool
	setTimeout(timeout time.Duration)
	setVersion(version string)
	setOut(file string)
	setDumpResult(file string)
	setExecutePath(file string)
}

//noinspection NonAsciiCharacters
type baseRunner struct {
	version    string
	command    string
	timeout    time.Duration
	out        string
	dumpResult string
	v8path     string
}

func (r baseRunner) setTimeout(timeout time.Duration) {
	r.timeout = timeout
}

func (r baseRunner) setVersion(version string) {
	r.version = version
}

func (r baseRunner) setExecutePath(file string) {
	r.v8path = file
}

func (r baseRunner) setOut(file string) {
	r.out = file
}
func (r baseRunner) setDumpResult(file string) {
	r.dumpResult = file
}

type Option func(*baseRunner)

const (
	commandDesigner       = "DESIGNER"
	commandCreateInfobase = "CREATEINFOBASE"
	commandEnterprise     = "ENTERPRISE"
)

const (
	EXITCODE_SUCCESS        = 0   // команда выполнена успешно.
	EXITCODE_FAIL           = 1   // при выполнении команды произошла ошибка.
	EXITCODE_DATABASE_ERROR = 101 // при выполнении команды обнаружены ошибки в данных.

)

var ERROR_CHECK_RUNNING = errors.New("error checking runeble")
var ERROR_VERSION_NOT_FOUND = errors.New("error Version is not found")
var ERROR_RUNNING_TIMEOUT = errors.New("error running Timeout executed")

var ERROR_RUNNING_FAIL = errors.New("error running v8 fail")
var ERROR_RUNNING_DATABASE_ERROR = errors.New("error running v8 database error")

func (conf *baseRunner) собратьПараметрыЗапуска() {

	//conf.args
	conf.args = []string{}

	conf.args = append(conf.args, string(conf.command))

	if conf.command == КомандаСоздатьБазу {
		// TODO Сделать замену /F на File= или /S на Server=
		log.Debugf("Выполняю замену </F> и </S> в строке <%s> на параметры для создания базы. ", conf.КлючСоединенияСБазой())
		conf.args = append(conf.args, strings.Replace(conf.КлючСоединенияСБазой(), "/F", "File=", 1))
	} else {
		conf.args = append(conf.args, conf.КлючСоединенияСБазой())
	}

	conf.добавитьАвторизацию()
	conf.добавитьКлючРазрешенияЗапуска()

	conf.args = append(conf.args, "/DisableStartupMessages")
	conf.args = append(conf.args, "/DisableStartupDialogs")

	conf.args = append(conf.args, "/AppAutoCheckVersion-")

	conf.args = append(conf.args, conf.пользовательскиеПараметрыЗапуска...)

	conf.добавитьВыводВФайл()

}

type RunebleOption func(options *RunOptions)

// private run func
const defaultFailedCode = 1

func WithTimeout(timeout time.Duration) RunebleOption {
	return func(r *RunOptions) {
		r.Timeout = timeout
	}
}

func WithOut(file string) RunebleOption {
	return func(r *RunOptions) {
		r.Out = file
	}
}

func WithDumpResult(file string) RunebleOption {
	return func(r *RunOptions) {
		r.DumpResult = file
	}
}

func WithVersion(version string) RunebleOption {
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

func RunWithOptions(runner runeble, options *RunOptions) (int, error) {

	var exitCode int

	if runner.check() {
		return defaultFailedCode, ERROR_CHECK_RUNNING
	}

	commandV8, err := getV8Path(options)

	if err != nil {
		return defaultFailedCode, err
	}

	// Create a new context and add a Timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	args := runner.args()
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
			exitCode = defaultFailedCode
		}
	} else {

		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()

	}

	// We want to check the context error to see if the Timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		return defaultFailedCode, ERROR_RUNNING_TIMEOUT
	}

	dumpCode := readDumpResult(options.DumpResult)

	switch dumpCode {

	case EXITCODE_SUCCESS:

		return EXITCODE_SUCCESS, nil

	case EXITCODE_FAIL:

		return EXITCODE_FAIL, ERROR_RUNNING_FAIL

	case EXITCODE_DATABASE_ERROR:

		return EXITCODE_DATABASE_ERROR, ERROR_RUNNING_DATABASE_ERROR

	default:

		return defaultFailedCode, errors.New("unknown error")
	}
}

func defaultOptions(command string) *RunOptions {

	var options RunOptions

	switch command {

	case commandDesigner:
	case commandEnterprise:
	case commandCreateInfobase:
	}

	return &options
}

func Run(runner runeble, opts ...RunebleOption) (runErr error) {

	options := defaultOptions(runner.command())

	for _, opt := range opts {
		opt(options)
	}

	return RunWithOptions(runner, options)

}

func (c *baseRunner) прочитатьФайлИнформации() (str string) {

	log.Debugf("Читаю файла информации 1С: %s", c.out)

	b, err := v8tools.ПрочитатьФайл1С(c.out) // just pass the file name
	if err != nil {
		log.Debugf("Обшибка чтения файла информации 1С %s: %v", c.out, err)
		str = ""
		return
		//fmt.Print(err)
	}

	str = string(b) // convert content to a 'string'

	return
}

func init() {

	//log.SetLevel(log.DebugLevel)

}

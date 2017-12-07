package config

import (
	"github.com/khorevaa/go-AutoUpdate1C/logging"

	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"runtime"
	"time"
)

// Profile contains environment specific options
type Config struct {
	Debug   bool `json:"debug"`
	V8      string
	log     *logging.Logger
	TimeOut time.Duration
}

type ConfigFn func() *Config

func (p *Config) Log() *logging.Logger {
	if p.log == nil {
		logger := logConfig()
		if p.Debug {
			logger.SetLevel(log.DebugLevel)
		}

		if runtime.GOOS == "windows" {
			logger.Formatter = &log.TextFormatter{ForceColors: true}
			logger.Out = colorable.NewColorableStdout()
		}

		p.log = logger
	}
	return p.log
}

func logConfig() *logging.Logger {

	return logging.New()

}

package config

import (
	"github.com/khorevaa/go-AutoUpdate1C/logging"

	log "github.com/sirupsen/logrus"
)

// Profile contains environment specific options
type Config struct {
	Debug bool `json:"debug"`
	V8    string
	log   *logging.Logger
}

func (p *Config) Log() *logging.Logger {
	if p.log == nil {
		logger := logConfig()
		if p.Debug {
			logger.SetLevel(log.DebugLevel)
		}

		p.log = logger
	}
	return p.log
}

func logConfig() *logging.Logger {

	return logging.New()

}

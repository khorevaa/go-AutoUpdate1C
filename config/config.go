package config

import (
	"github/Khorevaa/go-AutoUpdate1C/logging"

	log "github.com/sirupsen/logrus"
)

// Profile contains environment specific options
type Config struct {
	Master  string `json:"master"`
	Debug   bool   `json:"debug"`
	Restart bool   `json:"restart"`
	Sync    bool   `json:"sync"`
	V8      string
	log     *logging.Logger
}

func (p *Config) Log() *logging.Logger {
	if p.log == nil {
		logger := logConfig()
		logger.SetLevel(log.ErrorLevel)
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

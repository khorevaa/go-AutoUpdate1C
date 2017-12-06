package logging

import (
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	*log.Logger
}

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetLevel(log.ErrorLevel)
	//log.StandardLogger()
}

func New() *Logger {

	return &Logger{Logger: log.New()}
}

func (l *Logger) NewContextLogger(c log.Fields) func(f log.Fields) *log.Entry {
	return func(f log.Fields) *log.Entry {
		for k, v := range c {
			f[k] = v
		}
		return l.WithFields(f)
	}
}

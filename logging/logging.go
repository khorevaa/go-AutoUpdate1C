package logging

import (
	log "github.com/sirupsen/logrus"
)

type LogFunc func(f LogFeilds) Log
type Log struct {
	*log.Entry
}
type LogFeilds log.Fields

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

func (l *Logger) NewContextLogger(c LogFeilds) LogFunc {
	return func(f LogFeilds) Log {
		a := log.Fields{}
		for k, v := range c {
			a[k] = v
		}
		return l.WithFields(a)
	}
}

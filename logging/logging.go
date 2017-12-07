package logging

import (
	log "github.com/sirupsen/logrus"
)

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

	return &Logger{
		Logger: log.New(),
	}
}

func (l *Logger) NewContextLogger(c LogFeilds) Log {
	a := log.Fields{}
	for k, v := range c {
		a[k] = v
	}
	return Log{Entry: l.WithFields(a)}
}

func (l *Log) Context(c LogFeilds) Log {
	a := log.Fields{}
	for k, v := range c {
		a[k] = v
	}
	return Log{Entry: l.WithFields(a)}
}

func (l *Log) IfError(err error, m ...interface{}) {
	if err != nil {
		l.WithError(err).Error(m...)
	}
}

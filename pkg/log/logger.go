package log

import (
	stdlog "log"
)

type Logger interface {
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
}

type StdLogger struct{}

func New() *StdLogger {
	return new(StdLogger)
}

func (s *StdLogger) Errorf(format string, v ...interface{}) {
	stdlog.Printf(format, v...)
}

func (s *StdLogger) Infof(format string, v ...interface{}) {
	stdlog.Printf(format, v...)
}

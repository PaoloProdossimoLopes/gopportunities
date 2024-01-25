package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "::DEBUG::", logger.Flags()),
		info:    log.New(writer, "::INFO::", logger.Flags()),
		warning: log.New(writer, "::WARNING::", logger.Flags()),
		err:     log.New(writer, "::ERROR::", logger.Flags()),
		writer:  writer,
	}
}

func (l Logger) Debug(values ...interface{}) {
	l.debug.Println(values)
}

func (l Logger) Info(values ...interface{}) {
	l.info.Println(values)
}

func (l Logger) Warning(values ...interface{}) {
	l.warning.Println(values)
}

func (l Logger) Error(values ...interface{}) {
	l.err.Println(values)
}

func (l Logger) Debugf(format string, values ...interface{}) {
	l.debug.Printf(format, values)
}

func (l Logger) Infof(format string, values ...interface{}) {
	l.debug.Printf(format, values)
}

func (l Logger) Warningf(format string, values ...interface{}) {
	l.debug.Printf(format, values)
}

func (l Logger) Errorf(format string, values ...interface{}) {
	l.debug.Printf(format, values)
}

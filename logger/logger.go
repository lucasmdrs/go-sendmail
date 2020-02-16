package logger

import (
	"log"
	"os"
)

const (
	INFO int = iota
	WARNING
	ERROR
	DEBUG
)

var logLevels = map[string]int{
	"INFO":    INFO,
	"WARNING": WARNING,
	"ERROR":   ERROR,
	"DEBUG":   DEBUG,
}

type logger struct {
	level int
}

func DefaultLogger() Logger {
	if lvl, isSet := os.LookupEnv("LOG_LEVEL"); isSet {
		if logLevels[lvl] > INFO && logLevels[lvl] <= DEBUG {
			return &logger{level: logLevels[lvl]}
		}
	}
	return &logger{}
}

type Logger interface {
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
}

func (l logger) Info(args ...interface{}) {
	msg := append([]interface{}{"[INFO]"}, args...)
	l.log(INFO, msg...)
}
func (l logger) Warning(args ...interface{}) {
	msg := append([]interface{}{"[WARNING]"}, args...)
	l.log(WARNING, msg...)
}
func (l logger) Error(args ...interface{}) {
	msg := append([]interface{}{"[ERROR]"}, args...)
	l.log(ERROR, msg...)
}
func (l logger) Debug(args ...interface{}) {
	msg := append([]interface{}{"[DEBUG]"}, args...)
	l.log(DEBUG, msg...)
}

func (l logger) log(level int, args ...interface{}) {
	if l.level >= level {
		log.Println(args...)
	}
}

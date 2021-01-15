package logrus

import (
	log "github.com/sirupsen/logrus"
)

func NewUpperLogrusLogger() *log.Logger {
	l := log.New()
	std := log.StandardLogger()
	l.Level = std.Level
	l.Hooks = std.Hooks
	l.Formatter = std.Formatter
	l.Out = std.Out

	return l
}

package log4go

import (
	"github.com/lankoru/log4go/appenders"
	"github.com/lankoru/log4go/event"
	"github.com/lankoru/log4go/levels"
)

type Logger interface {
	Name() string
	Enabled(lvl levels.LogLevel) bool

	CreateSubLogger(name string) Logger
	Parent() Logger
	GetSubLogger(name string) Logger

	AddAppenders(a ...appenders.Appender)

	Log(ev event.LogEvent)

	WithFields(fields event.FieldsMap) *Ctx
	WithField(key string, value interface{}) *Ctx

	Fatalf(format string, args ...interface{})
	Critf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
}

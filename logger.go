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

	AddAppender(a appenders.Appender)

	Log(ev event.LogEvent)

	WithFields(fields event.FieldsMap) *Ctx
	WithField(key string, value interface{}) *Ctx

	Infof(format string, args ...interface{})
}

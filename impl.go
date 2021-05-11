package log4go

import (
	"sync"

	"github.com/lankoru/log4go/appenders"
	"github.com/lankoru/log4go/event"
	"github.com/lankoru/log4go/levels"
)

var _ Logger = (*logger)(nil)

type logger struct {
	mut sync.RWMutex

	name     string
	logLevel levels.LogLevel

	parent   Logger
	children []Logger

	appenders []appenders.Appender
}

func (l *logger) CreateSubLogger(name string) Logger {
	l.mut.RLock()
	defer l.mut.RUnlock()

	child := &logger{
		mut:       sync.RWMutex{},
		name:      name,
		parent:    l,
		appenders: l.appenders,
		logLevel:  l.logLevel,
		children:  make([]Logger, 0, 10),
	}

	l.children = append(l.children, child)

	return child
}

func (l *logger) GetSubLogger(name string) Logger {
	l.mut.RLock()
	defer l.mut.RUnlock()

	for i := range l.children {
		if l.children[i].Name() == name {
			return l.children[i]
		}
	}

	return nil
}

func (l *logger) Parent() Logger {
	l.mut.RLock()
	defer l.mut.RUnlock()

	return l.parent
}

func (l *logger) Name() string {
	l.mut.RLock()
	defer l.mut.RUnlock()

	return l.name
}

func (l *logger) Enabled(lvl levels.LogLevel) bool {
	l.mut.RLock()
	defer l.mut.RUnlock()

	return lvl <= l.logLevel
}

func (l *logger) AddAppender(a appenders.Appender) {
	l.mut.Lock()
	defer l.mut.Unlock()

	l.appenders = append(l.appenders, a)
}

func (l *logger) Log(e event.LogEvent) {
	if !l.Enabled(e.LogLevel) {
		return
	}

	l.mut.RLock()
	defer l.mut.RUnlock()

	e.TimeStamp = _clock.Now()
	e.Category = l.name

	for _, a := range l.appenders {
		a.Append(e)
	}
}

func (l *logger) WithField(key string, value interface{}) *Ctx {
	l.mut.RLock()
	defer l.mut.RUnlock()

	c := newCtx(l)
	return c.WithField(key, value)
}

func (l *logger) WithFields(fields event.FieldsMap) *Ctx {
	l.mut.RLock()
	defer l.mut.RUnlock()

	c := newCtx(l)
	return c.WithFields(fields)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.mut.RLock()
	defer l.mut.RUnlock()

	c := newCtx(l)
	c.Infof(format, args...)
}

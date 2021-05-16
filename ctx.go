package log4go

import (
	"fmt"
	"sync"

	"github.com/lankoru/log4go/event"
	"github.com/lankoru/log4go/levels"
)

type Ctx struct {
	mut sync.RWMutex

	logger Logger
	ev     event.LogEvent
}

func newCtx(l Logger) *Ctx {
	return &Ctx{
		mut:    sync.RWMutex{},
		logger: l,
		ev:     event.LogEvent{Fields: make(map[string]interface{})},
	}
}

func (c *Ctx) WithField(key string, value interface{}) *Ctx {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Fields[key] = value

	return c
}

func (c *Ctx) WithFields(fields event.FieldsMap) *Ctx {
	c.mut.Lock()
	defer c.mut.Unlock()

	for k, v := range fields {
		c.ev.Fields[k] = v
	}

	return c
}

func (c *Ctx) Fatalf(format string, args ...interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Message = fmt.Sprintf(format, args...)
	c.ev.LogLevel = levels.LevelFatal

	c.logger.Log(c.ev)
}

func (c *Ctx) Critf(format string, args ...interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Message = fmt.Sprintf(format, args...)
	c.ev.LogLevel = levels.LevelCritical

	c.logger.Log(c.ev)
}

func (c *Ctx) Errorf(format string, args ...interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Message = fmt.Sprintf(format, args...)
	c.ev.LogLevel = levels.LevelError

	c.logger.Log(c.ev)
}

func (c *Ctx) Warnf(format string, args ...interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Message = fmt.Sprintf(format, args...)
	c.ev.LogLevel = levels.LevelWarn

	c.logger.Log(c.ev)
}

func (c *Ctx) Infof(format string, args ...interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Message = fmt.Sprintf(format, args...)
	c.ev.LogLevel = levels.LevelInfo

	c.logger.Log(c.ev)
}

func (c *Ctx) Debugf(format string, args ...interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Message = fmt.Sprintf(format, args...)
	c.ev.LogLevel = levels.LevelDebug

	c.logger.Log(c.ev)
}

func (c *Ctx) Tracef(format string, args ...interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Message = fmt.Sprintf(format, args...)
	c.ev.LogLevel = levels.LevelTrace

	c.logger.Log(c.ev)
}

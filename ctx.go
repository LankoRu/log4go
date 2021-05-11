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

func (c *Ctx) Infof(format string, args ...interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ev.Message = fmt.Sprintf(format, args...)
	c.ev.LogLevel = levels.LevelInfo

	c.logger.Log(c.ev)
}

package log4go

import (
	"sync"

	"github.com/lankoru/log4go/levels"
)

const (
	defaultLogLevel = levels.LevelTrace
)

var _root = &logger{
	mut:       sync.RWMutex{},
	name:      "",
	parent:    nil,
	logLevel:  defaultLogLevel,
	appenders: nil,
	children:  make([]Logger, 0, 10),
}

func Root() Logger {
	return _root
}

func Reset() {
	_root.mut.Lock()
	defer _root.mut.Unlock()

	_root.logLevel = defaultLogLevel
	_root.children = make([]Logger, 0, 10)
	_root.appenders = nil
}

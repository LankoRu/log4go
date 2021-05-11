package appenders

import "github.com/lankoru/log4go/event"

type Appender interface {
	GetName() string

	Append(ev event.LogEvent)
}

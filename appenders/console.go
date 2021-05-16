package appenders

import (
	"os"

	"github.com/lankoru/log4go/event"
	"github.com/lankoru/log4go/layouts"
)

type consoleAppender struct {
	target *os.File

	layout layouts.Layout
}

func NewConsoleAppender(layout layouts.Layout) Appender {
	return &consoleAppender{
		target: os.Stdout,
		layout: layout,
	}
}

func (ca consoleAppender) GetName() string {
	return "console"
}

func (ca consoleAppender) Append(ev event.LogEvent) {
	_, _ = ca.target.Write(ca.layout.Format(ev))
}

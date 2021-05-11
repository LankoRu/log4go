package appenders

import (
	"os"

	"github.com/lankoru/log4go/event"
	"github.com/lankoru/log4go/layouts"
)

type ConsoleAppender struct {
	Layout layouts.Layout
	Target *os.File
}

func (cl ConsoleAppender) GetName() string {
	return "console"
}

func (cl ConsoleAppender) Append(ev event.LogEvent) {
	_, _ = cl.Target.Write(cl.Layout.Format(ev))
}

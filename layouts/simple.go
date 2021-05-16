package layouts

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/lankoru/log4go/event"
	"github.com/lankoru/log4go/levels"
)

type simpleLayout struct {
	params SimpleLayoutParams

	levelPainters map[levels.LogLevel]func(...interface{}) string
}

type SimpleLayoutParams struct {
	AutoColors bool
}

func NewSimpleLayout(params SimpleLayoutParams) Layout {
	return &simpleLayout{
		params: params,
		levelPainters: map[levels.LogLevel]func(...interface{}) string{
			levels.LevelFatal:    color.New(color.FgRed).SprintFunc(),
			levels.LevelCritical: color.New(color.FgRed).SprintFunc(),
			levels.LevelError:    color.New(color.FgRed).SprintFunc(),
			levels.LevelWarn:     color.New(color.FgYellow).SprintFunc(),
			levels.LevelInfo:     color.New(color.FgGreen).SprintFunc(),
			levels.LevelDebug:    color.New(color.FgBlue).SprintFunc(),
			levels.LevelTrace:    color.New(color.FgBlue).SprintFunc(),
		},
	}
}

func (sl simpleLayout) levelString(ev event.LogEvent) string {
	if !sl.params.AutoColors {
		return ev.LogLevel.String()
	}

	if f, ok := sl.levelPainters[ev.LogLevel]; ok {
		return f(ev.LogLevel.String())
	}

	return ev.LogLevel.String()
}

func (sl simpleLayout) Format(ev event.LogEvent) []byte {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%s - %s ", sl.levelString(ev), ev.Message))

	for k, v := range ev.Fields {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(fmt.Sprint(v))
		sb.WriteRune(' ')
	}

	msg := strings.TrimSpace(sb.String())

	return []byte(msg + string(newLine))
}

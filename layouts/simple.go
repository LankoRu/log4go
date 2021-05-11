package layouts

import (
	"fmt"
	"strings"

	"github.com/lankoru/log4go/event"
)

type simpleLayout struct {
}

func NewSimpleLayout() Layout {
	return &simpleLayout{}
}

func (sl simpleLayout) Format(ev event.LogEvent) []byte {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%s - %s ", ev.LogLevel.String(), ev.Message))

	for k, v := range ev.Fields {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(fmt.Sprint(v))
		sb.WriteRune(' ')
	}

	msg := strings.TrimSpace(sb.String())

	return []byte(msg + string(newLine))
}

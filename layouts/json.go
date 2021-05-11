package layouts

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/lankoru/log4go/event"
)

type jsonLayout struct {
}

func NewJSONLayout() Layout {
	return &jsonLayout{}
}

func (jl jsonLayout) Format(ev event.LogEvent) []byte {
	type jsonMessageModel struct {
		Msg       string                 `json:"msg"`
		Lvl       string                 `json:"lvl"`
		TimeStamp time.Time              `json:"timestamp"`
		Category  string                 `json:"category,omitempty"`
		Fields    map[string]interface{} `json:"fields,omitempty"`
	}

	bb, err := json.Marshal(&jsonMessageModel{
		Msg:       ev.Message,
		Lvl:       strings.ToLower(ev.LogLevel.String()),
		TimeStamp: ev.TimeStamp,
		Fields:    ev.Fields,
		Category:  ev.Category,
	})
	bb = append(bb, newLine...)
	if err != nil {
		fmt.Printf("json construction failed - %v\n", err)
		return nil
	}

	return bb
}

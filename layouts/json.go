package layouts

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/lankoru/log4go/event"
)

type jsonLayout struct {
	params JSONLayoutParams
}

type JSONLayoutParams struct {
	HideTimestamp bool
}

func NewJSONLayout(params JSONLayoutParams) Layout {
	return &jsonLayout{params: params}
}

func (jl jsonLayout) Format(ev event.LogEvent) []byte {
	type jsonMessageModel struct {
		Msg       string                 `json:"msg"`
		Lvl       string                 `json:"lvl"`
		TimeStamp *time.Time             `json:"timestamp,omitempty"`
		Category  string                 `json:"category,omitempty"`
		Fields    map[string]interface{} `json:"fields,omitempty"`
	}

	m := jsonMessageModel{
		Msg:      ev.Message,
		Lvl:      strings.ToLower(ev.LogLevel.String()),
		Fields:   ev.Fields,
		Category: ev.Category,
	}
	if !jl.params.HideTimestamp {
		m.TimeStamp = &ev.TimeStamp
	}

	bb, err := json.Marshal(&m)
	bb = append(bb, _newLine...)
	if err != nil {
		fmt.Printf("json construction failed - %v\n", err)
		return nil
	}

	return bb
}

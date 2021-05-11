package event

import (
	"time"

	"github.com/lankoru/log4go/levels"
)

type FieldsMap = map[string]interface{}

type LogEvent struct {
	Message   string
	Fields    FieldsMap
	LogLevel  levels.LogLevel
	TimeStamp time.Time
	Category  string
}

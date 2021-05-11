package levels

type LogLevel int

const (
	LevelFatal LogLevel = iota
	LevelCritical
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)

func (ll LogLevel) String() string {
	if s, ok := logLevelToStringMap[ll]; ok {
		return s
	}
	return "UNKNOWN"
}

var stringToLogLevelsMap = map[string]LogLevel{
	"TRACE": LevelTrace,
	"DEBUG": LevelDebug,
	"INFO":  LevelInfo,
	"WARN":  LevelWarn,
	"ERROR": LevelError,
	"CRIT":  LevelCritical,
	"FATAL": LevelFatal,
}

var logLevelToStringMap = map[LogLevel]string{
	LevelTrace:    "TRACE",
	LevelDebug:    "DEBUG",
	LevelInfo:     "INFO",
	LevelWarn:     "WARN",
	LevelError:    "ERROR",
	LevelCritical: "CRIT",
	LevelFatal:    "FATAL",
}

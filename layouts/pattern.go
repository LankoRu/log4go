package layouts

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"

	"github.com/lankoru/log4go/event"
	"github.com/lankoru/log4go/levels"
)

type patternLayout struct {
	params          PatterLayoutParams
	created         int64
	re              *regexp.Regexp
	levelPainters   map[levels.LogLevel]func(...interface{}) string
	categoryPainter func(...interface{}) string
}

type PatterLayoutParams struct {
	Pattern    string
	AutoColors bool
}

var defaultPatternTimeLayout = "2006-01-02 15:04:05.000000000 -0700 MST"

func NewPatternLayout(params PatterLayoutParams) Layout {
	return &patternLayout{
		params:  params,
		re:      regexp.MustCompile("%(\\w|%)(?:{([^}]+)})?"),
		created: time.Now().UnixNano(),
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

type caller struct {
	pc       uintptr
	file     string
	line     int
	ok       bool
	pkg      string
	fullpkg  string
	filename string
}

func getCaller() *caller {
	pc, file, line, ok := runtime.Caller(2)

	dir, fn := filepath.Split(file)
	bits := strings.Split(dir, "/")
	pkg := bits[len(bits)-2]

	if ok {
		return &caller{pc, file, line, ok, pkg, pkg, fn}
	}
	return nil
}

func (pl patternLayout) levelString(ev event.LogEvent) string {
	if !pl.params.AutoColors {
		return ev.LogLevel.String()
	}

	if f, ok := pl.levelPainters[ev.LogLevel]; ok {
		return f(ev.LogLevel.String())
	}

	return ev.LogLevel.String()
}

func (pl patternLayout) Format(ev event.LogEvent) []byte {
	cl := getCaller()
	r := ev.TimeStamp.UnixNano()

	msg := pl.re.ReplaceAllStringFunc(pl.params.Pattern, func(m string) string {
		parts := pl.re.FindStringSubmatch(m)
		switch parts[1] {
		case "c":
			return ev.Category
		case "C":
			return ev.Category
		case "d":
			if parts[2] == "" {
				return ev.TimeStamp.Format(defaultPatternTimeLayout)
			}
			return ev.TimeStamp.Format(parts[2])
		case "F":
			return cl.file
		case "l":
			return fmt.Sprintf("%s/%s:%d", cl.pkg, cl.filename, cl.line)
		case "L":
			return strconv.Itoa(cl.line)
		case "m":
			return fmt.Sprintf(ev.Message)
		case "n":
			// skip new line separator, we will append it in the end
			return ""
		case "p":
			return pl.levelString(ev)
		case "r":
			return strconv.FormatInt((r-pl.created)/100000, 10)
		case "x":
			sb := strings.Builder{}
			for k, v := range ev.Fields {
				sb.WriteString(k)
				sb.WriteString("=")
				sb.WriteString(fmt.Sprint(v))
				sb.WriteRune(' ')
			}
			return strings.TrimSpace(sb.String())
		case "X":
			return ""
		case "%":
			return "%"
		}
		return m
	})

	return []byte(strings.TrimSpace(msg) + string(newLine))
}

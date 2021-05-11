package layouts

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/lankoru/log4go/event"
)

type patternLayout struct {
	Pattern string
	created int64
	re      *regexp.Regexp
}

var defaultPatternTimeLayout = "2006-01-02 15:04:05.000000000 -0700 MST"

func NewPatternLayout(pattern string) Layout {
	return &patternLayout{
		Pattern: pattern,
		re:      regexp.MustCompile("%(\\w|%)(?:{([^}]+)})?"),
		created: time.Now().UnixNano(),
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

func (pl patternLayout) Format(ev event.LogEvent) []byte {
	cl := getCaller()
	r := ev.TimeStamp.UnixNano()

	msg := pl.re.ReplaceAllStringFunc(pl.Pattern, func(m string) string {
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
			return ev.LogLevel.String()
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

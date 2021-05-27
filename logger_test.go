package log4go_test

import (
	"time"

	"github.com/lankoru/log4go"
	"github.com/lankoru/log4go/appenders"
	"github.com/lankoru/log4go/layouts"
)

type testClock struct {
}

func (t testClock) Now() time.Time {
	return time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
}

func runOperations() {
	log4go.Root().Critf("message %d", 1)

	log4go.Root().WithField("key", "value").Errorf("message %d", 2)

	sub := log4go.Root().CreateSubLogger("sub")
	sub.WithField("key", "value").Warnf("message %d", 3)

	sub = log4go.Root().GetSubLogger("sub")
	sub.WithField("key", "value").Infof("message %d", 4)
	sub.WithField("key", "value").Debugf("message %d", 5)
	sub.WithField("key", "value").Tracef("message %d", 6)
}

func ExampleNewSimpleLayout() {
	defer log4go.Reset()

	log4go.Root().AddAppenders(
		appenders.NewConsoleAppender(layouts.NewSimpleLayout(layouts.SimpleLayoutParams{AutoColors: true})),
	)

	runOperations()

	// Output:
	// CRIT - message 1
	// ERROR - message 2 key=value
	// WARN - message 3 key=value
	// INFO - message 4 key=value
	// DEBUG - message 5 key=value
	// TRACE - message 6 key=value
}

func ExampleNewJSONLayout() {
	defer log4go.Reset()

	log4go.MockClock(testClock{})
	defer log4go.MockClock(nil)

	log4go.Root().AddAppenders(
		appenders.NewConsoleAppender(layouts.NewJSONLayout(layouts.JSONLayoutParams{HideTimestamp: false})),
		appenders.NewConsoleAppender(layouts.NewJSONLayout(layouts.JSONLayoutParams{HideTimestamp: true})),
	)

	runOperations()

	// Output:
	// {"msg":"message 1","lvl":"crit","timestamp":"2020-01-01T00:00:00Z"}
	// {"msg":"message 1","lvl":"crit"}
	// {"msg":"message 2","lvl":"error","timestamp":"2020-01-01T00:00:00Z","fields":{"key":"value"}}
	// {"msg":"message 2","lvl":"error","fields":{"key":"value"}}
	// {"msg":"message 3","lvl":"warn","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// {"msg":"message 3","lvl":"warn","category":"sub","fields":{"key":"value"}}
	// {"msg":"message 4","lvl":"info","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// {"msg":"message 4","lvl":"info","category":"sub","fields":{"key":"value"}}
	// {"msg":"message 5","lvl":"debug","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// {"msg":"message 5","lvl":"debug","category":"sub","fields":{"key":"value"}}
	// {"msg":"message 6","lvl":"trace","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// {"msg":"message 6","lvl":"trace","category":"sub","fields":{"key":"value"}}
}

func ExampleNewPatternLayout() {
	defer log4go.Reset()

	log4go.MockClock(testClock{})
	defer log4go.MockClock(nil)

	log4go.Root().AddAppenders(
		appenders.NewConsoleAppender(layouts.NewPatternLayout(layouts.PatterLayoutParams{
			Pattern:    "%d{2006-01-02T15:04:05.000-07:00MST} %p [%c] %m %x",
			AutoColors: true,
		})))

	runOperations()

	// Output:
	// 2020-01-01T00:00:00.000+00:00UTC CRIT [] message 1
	// 2020-01-01T00:00:00.000+00:00UTC ERROR [] message 2 key=value
	// 2020-01-01T00:00:00.000+00:00UTC WARN [sub] message 3 key=value
	// 2020-01-01T00:00:00.000+00:00UTC INFO [sub] message 4 key=value
	// 2020-01-01T00:00:00.000+00:00UTC DEBUG [sub] message 5 key=value
	// 2020-01-01T00:00:00.000+00:00UTC TRACE [sub] message 6 key=value
}

func ExampleAllLayouts() {
	defer log4go.Reset()

	log4go.MockClock(testClock{})
	defer log4go.MockClock(nil)

	log4go.Root().AddAppenders(
		appenders.NewConsoleAppender(layouts.NewSimpleLayout(layouts.SimpleLayoutParams{AutoColors: true})),
		appenders.NewConsoleAppender(layouts.NewJSONLayout(layouts.JSONLayoutParams{})),
	)

	runOperations()

	// Output:
	// CRIT - message 1
	// {"msg":"message 1","lvl":"crit","timestamp":"2020-01-01T00:00:00Z"}
	// ERROR - message 2 key=value
	// {"msg":"message 2","lvl":"error","timestamp":"2020-01-01T00:00:00Z","fields":{"key":"value"}}
	// WARN - message 3 key=value
	// {"msg":"message 3","lvl":"warn","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// INFO - message 4 key=value
	// {"msg":"message 4","lvl":"info","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// DEBUG - message 5 key=value
	// {"msg":"message 5","lvl":"debug","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// TRACE - message 6 key=value
	// {"msg":"message 6","lvl":"trace","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
}

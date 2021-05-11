package log4go_test

import (
	"os"
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
	log4go.Root().Infof("message %d", 1)

	log4go.Root().WithField("key", "value").Infof("message %d", 2)

	sub := log4go.Root().CreateSubLogger("sub")
	sub.WithField("key", "value").Infof("message %d", 3)

	sub = log4go.Root().GetSubLogger("sub")
	sub.WithField("key", "value").Infof("message %d", 4)
}

func ExampleNewSimpleLayout() {
	defer log4go.Reset()

	log4go.Root().AddAppender(&appenders.ConsoleAppender{
		Layout: layouts.NewSimpleLayout(),
		Target: os.Stdout,
	})

	runOperations()

	// Output:
	// INFO - message 1
	// INFO - message 2 key=value
	// INFO - message 3 key=value
	// INFO - message 4 key=value
}

func ExampleNewJSONLayout() {
	defer log4go.Reset()

	log4go.MockClock(testClock{})
	defer log4go.MockClock(nil)

	log4go.Root().AddAppender(&appenders.ConsoleAppender{
		Layout: layouts.NewJSONLayout(),
		Target: os.Stdout,
	})

	runOperations()

	// Output:
	// {"msg":"message 1","lvl":"info","timestamp":"2020-01-01T00:00:00Z"}
	// {"msg":"message 2","lvl":"info","timestamp":"2020-01-01T00:00:00Z","fields":{"key":"value"}}
	// {"msg":"message 3","lvl":"info","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// {"msg":"message 4","lvl":"info","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
}

func ExampleNewPatternLayout() {
	defer log4go.Reset()

	log4go.MockClock(testClock{})
	defer log4go.MockClock(nil)

	log4go.Root().AddAppender(&appenders.ConsoleAppender{
		Layout: layouts.NewPatternLayout("%d{2006-01-02T15:04:05.000-07:00MST} %p [%c] %m %x"),
		Target: os.Stdout,
	})

	runOperations()

	// Output:
	// 2020-01-01T00:00:00.000+00:00UTC INFO [] message 1
	// 2020-01-01T00:00:00.000+00:00UTC INFO [] message 2 key=value
	// 2020-01-01T00:00:00.000+00:00UTC INFO [sub] message 3 key=value
	// 2020-01-01T00:00:00.000+00:00UTC INFO [sub] message 4 key=value
}

func ExampleAllLayouts() {
	defer log4go.Reset()

	log4go.MockClock(testClock{})
	defer log4go.MockClock(nil)

	log4go.Root().AddAppender(&appenders.ConsoleAppender{
		Layout: layouts.NewSimpleLayout(),
		Target: os.Stdout,
	})
	log4go.Root().AddAppender(&appenders.ConsoleAppender{
		Layout: layouts.NewJSONLayout(),
		Target: os.Stdout,
	})

	runOperations()

	// Output:
	// INFO - message 1
	// {"msg":"message 1","lvl":"info","timestamp":"2020-01-01T00:00:00Z"}
	// INFO - message 2 key=value
	// {"msg":"message 2","lvl":"info","timestamp":"2020-01-01T00:00:00Z","fields":{"key":"value"}}
	// INFO - message 3 key=value
	// {"msg":"message 3","lvl":"info","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
	// INFO - message 4 key=value
	// {"msg":"message 4","lvl":"info","timestamp":"2020-01-01T00:00:00Z","category":"sub","fields":{"key":"value"}}
}

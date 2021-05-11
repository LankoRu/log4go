package log4go

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time {
	return time.Now()
}

var _clock Clock = realClock{}

func MockClock(c Clock) {
	if c == nil {
		_clock = realClock{}
	} else {
		_clock = c
	}
}

package layouts

import "github.com/lankoru/log4go/event"

type Layout interface {
	Format(ev event.LogEvent) []byte
}

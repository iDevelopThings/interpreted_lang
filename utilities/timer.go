package utilities

import (
	"time"

	"arc/log"
)

type Timer struct {
	// The name of the timer.
	name string
	// The time at which the timer was started.
	start time.Time
	// The time at which the timer was stopped.
	stop time.Time
}

func NewTimer(name string) *Timer {
	t := &Timer{
		name:  name,
		start: time.Now(),
	}

	return t
}

func (t *Timer) Stop() {
	t.stop = time.Now()
}

func (t *Timer) String() string {
	return t.name + ": " + t.stop.Sub(t.start).String()
}

func (t *Timer) StopAndLog() {
	t.Stop()
	log.Helper()
	l := log.GetLevel()
	if l != log.DebugLevel {
		log.SetLevel(log.DebugLevel)
		defer log.SetLevel(l)
	}
	log.Debugf(t.String())
}

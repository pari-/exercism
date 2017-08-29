// Package clock provides the implementation of exercism's "clock" exercise.
package clock

import (
	"fmt"
)

// testVersion
const testVersion = 4

type Clock struct {
	minute, hour int
}

// New creates a new Clock set to a specific time.
func New(hour, minute int) Clock {
	mTotal := (minute + (hour * 60)) % (60 * 24)
	if mTotal < 0 {
		mTotal += (60 * 24)
	}
	clock := Clock{minute: (mTotal % 60), hour: (mTotal / 60)}
	return clock
}

// String prints a Clock's time in the format "hh:mm".
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add adds m minutes to a specified Clock.
func (clock Clock) Add(minutes int) Clock {
	return New(clock.hour, int(clock.minute)+minutes)
}

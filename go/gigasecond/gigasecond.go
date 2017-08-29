// Package gigasecond provides the implementation of exercism's "gigasecond" exercise.
package gigasecond

// import path for needed packages
import (
	"time"
)

// consts
const testVersion = 4

// AddGigasecond adds a Gigasecond to a specific time.
func AddGigasecond(before time.Time) time.Time {
	gigaNanoSecond := time.Duration(1000000000) * time.Second
	return before.Add(gigaNanoSecond)
}

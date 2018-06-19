// Package gigasecond computes the moment when someone lived for 10^9 seconds
package gigasecond

import (
	"time"
)

// AddGigasecond Returns the moment when someone lived for 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}

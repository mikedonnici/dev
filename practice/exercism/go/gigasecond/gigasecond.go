// Package gigasecond calculates the moment when someone has lived for 10^9 seconds
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to the input time t
func AddGigasecond(t time.Time) time.Time {

	return t.Add(1e9 * time.Second)
}
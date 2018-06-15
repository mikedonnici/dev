// Package clock implements a clock
package clock

import (
	"fmt"
	"time"
)

// referenceEpoch is set to midnight on Jan 1, 2000 (arbitrarily).
// The date itself is not important and is used as a reference point to wind clocks back or forward.
const referenceEpoch = 946684800

type clock struct {
	h int
	m int
}

// New returns a value of type clock
func New(h, m int) clock {
	c := clock{
		h: h,
		m: m,
	}
	return c.set()
}

// Add adds a number of minutes to a clock value
func (c clock) Add(minutes int) clock {
	ut := c.referenceEpoch() + int64(minutes*60)
	t := time.Unix(ut, 0).UTC()
	c.h = t.Hour()
	c.m = t.Minute()
	return c
}

// String returns the clock time as a string value
func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// set the clock
func (c clock) set() clock {
	t := c.referenceTime()
	c.h = t.Hour()
	c.m = t.Minute()
	return c
}

func (c clock) referenceTime() time.Time {
	return time.Unix(c.referenceEpoch(), 0).UTC()
}

func (c clock) referenceEpoch() int64 {
	return int64(referenceEpoch + (c.h * 3600) + (c.m * 60))
}

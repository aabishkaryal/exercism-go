// Package gigasecond implements functions for gigasecond addition
package gigasecond

import (
	"math"
	"time"
)

// Function AddGigasecond adds one gigasecond to given time and returns the new time
func AddGigasecond(t time.Time) time.Time {
	// 10^9 seconds is one gigasecond
	gigaSecond := int(math.Pow10(9)) * int(time.Second)
	t = t.Add(time.Duration(gigaSecond))
	return t
}

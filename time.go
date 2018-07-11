package senml

import (
	"math"
	"time"
)

// Time converts a Go time.Time to a SenML time (seconds since the epoch as a floating point number).
func Time(t time.Time) float64 {
	return float64(t.UnixNano()) / float64(time.Second)
}

// GoTime converts a SenML time to a Go time.Time.
func GoTime(f float64) time.Time {
	sec := math.Trunc(f)
	nsec := f*float64(time.Second) - sec*float64(time.Second)
	return time.Unix(int64(sec), int64(nsec))
}

var maxRelativeTime = math.Pow(2, 28)

func absoluteTime(t float64, ref float64) float64 {
	if t <= maxRelativeTime {
		return t + ref
	}
	return t
}

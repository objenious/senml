package senml

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t0 := time.Date(2018, 07, 11, 0, 0, 0, 100000000, time.UTC)
	st := Time(t0)
	if !GoTime(st).Equal(t0) {
		t.Errorf("Time(GoTime) does not return the right time")
	}
}

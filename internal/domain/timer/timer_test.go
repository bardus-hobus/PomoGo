package timer

import (
	"testing"
	"time"
)

func TestTimerDuration(t *testing.T) {
	d := 25 * time.Minute
	tmr := New(d)
	if tmr.Duration() != d {
		t.Fatalf("expected duration %v, got %v", d, tmr.Duration())
	}
}

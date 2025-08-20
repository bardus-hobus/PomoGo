package timer

import "time"

// Timer represents a pomodoro timer.
type Timer struct {
	duration time.Duration
}

// New creates a new Timer.
func New(duration time.Duration) Timer {
	return Timer{duration: duration}
}

// Duration returns the timer duration.
func (t Timer) Duration() time.Duration {
	return t.duration
}

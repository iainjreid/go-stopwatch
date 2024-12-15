// Package stopwatch offers a simple, no nonsence implementation of a stopwatch.
package stopwatch

import (
	"time"
)

type Stopwatch struct {
	Name      string
	StartedAt time.Time
	StoppedAt time.Time
	Splits    map[string]time.Time
}

// NewStopwatch creates a new stopwatch struct with the provided name.
func NewStopwatch(name string) *Stopwatch {
	return &Stopwatch{
		Name:   name,
		Splits: map[string]time.Time{},
	}
}

// Start starts the stopwatch.
func (s *Stopwatch) Start() {
	s.StartedAt = time.Now()
}

// Stop stops the stopwatch
func (s *Stopwatch) Stop() {
	s.StoppedAt = time.Now()
}

// Split records a named split
func (s *Stopwatch) Split(name string) {
	s.Splits[name] = time.Now()
}

// Reset clears the stopwatch
func (s *Stopwatch) Reset() {
	s.StartedAt = time.Time{}
	s.StoppedAt = time.Time{}
	s.Splits = map[string]time.Time{}
}

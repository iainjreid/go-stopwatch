package stopwatch

import (
	"testing"
)

// TestStart calls sw.Start, ensuring that the StartedAt property is correctly
// handled.
func TestStart(t *testing.T) {
	testStart(t, NewStopwatch("GladysStopwatch"))
}

func testStart(t *testing.T, sw *Stopwatch) {
	if !sw.StartedAt.IsZero() {
		t.Fatal("#StartedAt should be initiallised with a zero time value")
	}

	sw.Start()
	if sw.StartedAt.IsZero() {
		t.Fatal("Start() should set #StartedAt to a non-zero time value")
	}
}

// TestStop calls sw.Stop, ensuring that the StoppedAt property is correctly
// handled.
func TestStop(t *testing.T) {
	testStop(t, NewStopwatch("GladysStopwatch"))
}

func testStop(t *testing.T, sw *Stopwatch) {
	if !sw.StoppedAt.IsZero() {
		t.Fatal("#StoppedAt should be initiallised with a zero time value")
	}

	sw.Stop()
	if sw.StoppedAt.IsZero() {
		t.Fatal("Stop() should set #StoppedAt to a non-zero time value")
	}
}

// TestSplit calls sw.Split, ensuring that the Splits are properly added.
func TestSplit(t *testing.T) {
	testSplit(t, NewStopwatch("GladysStopwatch"))
}

func testSplit(t *testing.T, sw *Stopwatch) {
	if len(sw.Splits) != 0 {
		t.Fatal("#Splits should be initiallised with an empty value")
	}

	sw.Split("GladysLap")
	if len(sw.Splits) != 1 {
		t.Fatal("#Splits should have a length of 1")
	}

	split, ok := sw.Splits["GladysLap"]

	if !ok {
		t.Fatal("Split() should record a named split")
	}

	if split.IsZero() {
		t.Fatal("Split() should record a non-zero time value")
	}
}

// TestReset ensures that sw.Reset properyl resets all values
func TestReset(t *testing.T) {
	sw := NewStopwatch("GladysStopwatch")

	testStart(t, sw)
	testSplit(t, sw)
	testStop(t, sw)

	sw.Reset()

	if !sw.StartedAt.IsZero() {
		t.Fatal("#StartedAt should have been reset to a zero time value")
	}

	if !sw.StoppedAt.IsZero() {
		t.Fatal("#StoppedAt should have been reset to a zero time value")
	}

	if len(sw.Splits) != 0 {
		t.Fatal("#Splits should have been reset with an empty value")
	}
}

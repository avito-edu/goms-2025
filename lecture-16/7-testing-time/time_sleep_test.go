package testing_time

import (
	"reflect"
	"testing"
	"time"
)

type Sleeper interface {
	Sleep(time.Duration)
}

type RealSleeper struct{}

func (RealSleeper) Sleep(d time.Duration) {
	time.Sleep(d)
}

func Countdown(sleeper Sleeper, n int) {
	for i := n; i > 0; i-- {
		sleeper.Sleep(1 * time.Second)
	}
}

type SpySleeper struct {
	calls     int
	durations []time.Duration
}

func (s *SpySleeper) Sleep(d time.Duration) {
	s.calls++
	s.durations = append(s.durations, d)
}

func TestCountdown(t *testing.T) {
	sleeper := &SpySleeper{}
	Countdown(sleeper, 3)

	if sleeper.calls != 3 {
		t.Errorf("expected 3 sleeps, got %d", sleeper.calls)
	}

	expected := []time.Duration{1 * time.Second, 1 * time.Second, 1 * time.Second}
	if !reflect.DeepEqual(sleeper.durations, expected) {
		t.Errorf("expected sleeps %v, got %v", expected, sleeper.durations)
	}
}

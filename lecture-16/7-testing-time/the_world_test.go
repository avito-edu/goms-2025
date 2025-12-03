package testing_time

import (
	"testing"
	"time"
)

func IsMorningBad() bool {
	hour := time.Now().Hour()
	return hour >= 5 && hour < 12
}

func TestIsMorningBad(t *testing.T) {
	if !IsMorningBad() {
		t.Error("Expected morning time")
	}
}

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time {
	return time.Now()
}

func IsMorning(clock Clock) bool {
	hour := clock.Now().Hour()
	return hour >= 5 && hour < 12
}

type MockClock struct {
	fixedTime time.Time
}

func (m MockClock) Now() time.Time {
	return m.fixedTime
}

func TestIsMorning(t *testing.T) {
	tests := []struct {
		time time.Time
		want bool
	}{
		{time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC), true},   // утро
		{time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC), false}, // день
	}

	for _, tt := range tests {
		clock := MockClock{fixedTime: tt.time}
		got := IsMorning(clock)
		if got != tt.want {
			t.Errorf("IsMorning(%v) = %v, want %v", tt.time, got, tt.want)
		}
	}
}

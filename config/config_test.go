package config

import (
	"testing"
	"time"
)

func TestTimeoutDuration(t *testing.T) {
	tests := []struct {
		name     string
		timeout  int
		expected time.Duration
	}{
		{"Zero timeout", 0, 0 * time.Second},
		{"Positive timeout", 30, 30 * time.Second},
		{"Large timeout", 300, 300 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := RodConfig{Timeout: tt.timeout}
			if got := rc.TimeoutDuration(); got != tt.expected {
				t.Errorf("TimeoutDuration() = %v, want %v", got, tt.expected)
			}
		})
	}
}

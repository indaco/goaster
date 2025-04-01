package viewmodel

import "testing"

func TestGetToastEntranceClass(t *testing.T) {
	tests := []struct {
		position string
		expected string
	}{
		{"top-left", "gttShowFromTop"},
		{"top-center", "gttShowFromTop"},
		{"bottom-right", "gttShowFromBottom"},
		{"unknown", "gttShowFromBottom"},
	}

	for _, tt := range tests {
		if got := GetToastEntranceClass(tt.position); got != tt.expected {
			t.Errorf("for position %q, expected %q, got %q", tt.position, tt.expected, got)
		}
	}
}

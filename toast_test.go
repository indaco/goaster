package goaster

import (
	"testing"
)

func TestNewToast(t *testing.T) {
	tests := []struct {
		name          string
		message       string
		level         Level
		expectedMsg   string
		expectedLevel Level
	}{
		{
			name:          "info toast",
			message:       "Something happened",
			level:         InfoLevel,
			expectedMsg:   "Something happened",
			expectedLevel: InfoLevel,
		},
		{
			name:          "error toast",
			message:       "Something went wrong",
			level:         ErrorLevel,
			expectedMsg:   "Something went wrong",
			expectedLevel: ErrorLevel,
		},
		{
			name:          "success toast",
			message:       "All good",
			level:         SuccessLevel,
			expectedMsg:   "All good",
			expectedLevel: SuccessLevel,
		},
		{
			name:          "warning toast",
			message:       "Be careful",
			level:         WarningLevel,
			expectedMsg:   "Be careful",
			expectedLevel: WarningLevel,
		},
		{
			name:          "default toast",
			message:       "Default message",
			level:         DefaultLevel,
			expectedMsg:   "Default message",
			expectedLevel: DefaultLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toast := NewToast(tt.message, tt.level)
			if toast.Message != tt.expectedMsg {
				t.Errorf("Message: expected %q, got %q", tt.expectedMsg, toast.Message)
			}
			if toast.Level != tt.expectedLevel {
				t.Errorf("Level: expected %q, got %q", tt.expectedLevel, toast.Level)
			}
		})
	}
}

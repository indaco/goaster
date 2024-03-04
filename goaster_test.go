package goaster

import (
	"github.com/a-h/templ"
	"reflect"
	"testing"
)

// TestNewToasterDefault checks if a new Toaster is created with default values.
func TestNewToasterDefault(t *testing.T) {
	toaster := NewToaster()

	if toaster.Border != true {
		t.Errorf("expected Border to be true, got %v", toaster.Border)
	}

	if toaster.ShowIcon != true {
		t.Errorf("expected ShowIcon to be true, got %v", toaster.ShowIcon)
	}

	if toaster.Button != true {
		t.Errorf("expected Button to be true, got %v", toaster.Button)
	}

	if toaster.AutoDismiss != true {
		t.Errorf("expected AutoDismiss to be true, got %v", toaster.AutoDismiss)
	}

	if toaster.Position != BottomRight {
		t.Errorf("expected Position to be BottomRight, got %v", toaster.Position)
	}
}

// TestToasterOptions uses table-driven tests to verify the functional options of the Toaster.
func TestToasterOptions(t *testing.T) {
	tests := []struct {
		name     string
		options  []Option
		expected Toaster
	}{
		{
			name:    "Default options",
			options: nil,
			expected: Toaster{
				Border:      true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithBorder false",
			options: []Option{
				WithBorder(false),
			},
			expected: Toaster{
				Border:      false,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithShowIcon false",
			options: []Option{
				WithShowIcon(false),
			},
			expected: Toaster{
				Border:      true,
				ShowIcon:    false,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithButton false",
			options: []Option{
				WithButton(false),
			},
			expected: Toaster{
				Border:      true,
				ShowIcon:    true,
				Button:      false,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithAutoDismiss false",
			options: []Option{
				WithAutoDismiss(false),
			},
			expected: Toaster{
				Border:      true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: false,
				Animation:   true,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithAnimation false",
			options: []Option{
				WithAnimation(false),
			},
			expected: Toaster{
				Border:      true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   false,
				ProgressBar: true,
				Position:    BottomRight,
			},
		},
		{
			name: "WithProgressBar false",
			options: []Option{
				WithProgressBar(false),
			},
			expected: Toaster{
				Border:      true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Animation:   true,
				ProgressBar: false,
				Position:    BottomRight,
			},
		},
		{
			name: "WithPosition TopLeft",
			options: []Option{
				WithPosition(TopLeft),
			},
			expected: Toaster{
				Border:      true,
				ShowIcon:    true,
				Button:      true,
				AutoDismiss: true,
				Position:    TopLeft,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toaster := NewToaster(tt.options...)

			// Check each field in expected Toaster struct
			if toaster.Border != tt.expected.Border {
				t.Errorf("Border: expected %v, got %v", tt.expected.Border, toaster.Border)
			}
			if toaster.ShowIcon != tt.expected.ShowIcon {
				t.Errorf("ShowIcon: expected %v, got %v", tt.expected.ShowIcon, toaster.ShowIcon)
			}
			if toaster.Button != tt.expected.Button {
				t.Errorf("Button: expected %v, got %v", tt.expected.Button, toaster.Button)
			}
			if toaster.AutoDismiss != tt.expected.AutoDismiss {
				t.Errorf("AutoDismiss: expected %v, got %v", tt.expected.AutoDismiss, toaster.AutoDismiss)
			}
			if toaster.Position != tt.expected.Position {
				t.Errorf("Position: expected %v, got %v", tt.expected.Position, toaster.Position)
			}
		})
	}
}

// TestWithIcon verifies the WithIcon functional option for custom icons.
func TestWithIcon(t *testing.T) {
	customIcon := "<svg>Custom Icon</svg>"
	toaster := NewToaster(WithIcon(SuccessLevel, customIcon))

	if toaster.Icons[SuccessLevel] != customIcon {
		t.Errorf("WithIcon(SuccessLevel, customIcon) failed; expected custom icon, got %v", toaster.Icons[SuccessLevel])
	}
}

// TestQueueAdd verifies adding messages to the queue.
func TestQueueAdd(t *testing.T) {
	toaster := NewToaster()
	toaster.PushSuccess("Success message")
	toaster.PushError("Error message")

	if toaster.queue.Size() != 2 {
		t.Fatalf("expected 2 messages in the queue, got %d", toaster.queue.Size())
	}

	if toaster.queue.elements[0].Level != SuccessLevel || toaster.queue.elements[1].Level != ErrorLevel {
		t.Errorf("queue does not contain the expected messages in the correct order")
	}
}

func TestAddMessageMethods(t *testing.T) {
	tests := []struct {
		name       string
		addMessage func(t *Toaster) // Function to execute adding the message
		expected   Toast            // Expected message in the queue
	}{
		{
			name: "PushSuccess",
			addMessage: func(t *Toaster) {
				t.PushSuccess("Success message")
			},
			expected: Toast{
				Message: "Success message",
				Level:   SuccessLevel,
			},
		},
		{
			name: "PushError",
			addMessage: func(t *Toaster) {
				t.PushError("Error message")
			},
			expected: Toast{
				Message: "Error message",
				Level:   ErrorLevel,
			},
		},
		{
			name: "PushWarning",
			addMessage: func(t *Toaster) {
				t.PushWarning("Warning message")
			},
			expected: Toast{
				Message: "Warning message",
				Level:   WarningLevel,
			},
		},
		{
			name: "PushInfo",
			addMessage: func(t *Toaster) {
				t.PushInfo("Info message")
			},
			expected: Toast{
				Message: "Info message",
				Level:   InfoLevel,
			},
		},
		{
			name: "PushDefault",
			addMessage: func(t *Toaster) {
				t.PushDefault("Default message")
			},
			expected: Toast{
				Message: "Default message",
				Level:   DefaultLevel,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toaster := NewToaster() // Create a new Toaster instance for each test
			tt.addMessage(toaster)  // Execute the method to add a message

			// Check if there's exactly one message in the queue
			if toaster.queue.Size() != 1 {
				t.Fatalf("expected 1 message in the queue, got %d", toaster.queue.Size())
			}

			// Check if the message in the queue matches the expected message
			got := toaster.queue.elements[0]
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected message %+v, got %+v", tt.expected, got)
			}

			// Clear the queue for the next test
			toaster.queue.elements = []Toast{}
		})
	}
}

func TestToasterNotificationMethods(t *testing.T) {
	tests := []struct {
		name          string
		methodToTest  func(*Toaster, string) templ.Component
		expectedLevel Level
		message       string
	}{
		{
			name:          "Default notification",
			methodToTest:  (*Toaster).Default,
			expectedLevel: DefaultLevel,
			message:       "Default message",
		},
		{
			name:          "Success notification",
			methodToTest:  (*Toaster).Success,
			expectedLevel: SuccessLevel,
			message:       "Success message",
		},
		{
			name:          "Error notification",
			methodToTest:  (*Toaster).Error,
			expectedLevel: ErrorLevel,
			message:       "Error message",
		},
		{
			name:          "Warning notification",
			methodToTest:  (*Toaster).Warning,
			expectedLevel: WarningLevel,
			message:       "Warning message",
		},
		{
			name:          "Info notification",
			methodToTest:  (*Toaster).Info,
			expectedLevel: InfoLevel,
			message:       "Info message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toaster := NewToaster() // Assuming NewToaster correctly initializes the queue
			tt.methodToTest(toaster, tt.message)

			// Validate the queue has exactly one message with the expected content and level
			if toaster.queue.Size() != 1 {
				t.Fatalf("expected 0 message in the queue, got %d", toaster.queue.Size())
			}
		})
	}
}

package goaster

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestToasterDefault(t *testing.T) {
	toaster := ToasterDefaults()

	if toaster.Variant != "" {
		t.Errorf("expected Variant to be an empty string, got %v", toaster.Variant)
	}

	if toaster.Border != true {
		t.Errorf("expected Border to be true, got %v", toaster.Border)
	}

	if toaster.Rounded != true {
		t.Errorf("expected Rounded to be true, got %v", toaster.Rounded)
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

// TestQueueAdd verifies adding messages to the queue.
func TestQueueAdd(t *testing.T) {
	toaster := ToasterDefaults()
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
			toaster := ToasterDefaults() // Create a new Toaster instance for each test
			tt.addMessage(toaster)       // Execute the method to add a message

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
		{"Default", (*Toaster).Default, DefaultLevel, "Default message"},
		{"Success", (*Toaster).Success, SuccessLevel, "Success message"},
		{"Error", (*Toaster).Error, ErrorLevel, "Error message"},
		{"Warning", (*Toaster).Warning, WarningLevel, "Warning message"},
		{"Info", (*Toaster).Info, InfoLevel, "Info message"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toaster := ToasterDefaults()
			component := tt.methodToTest(toaster, tt.message)

			var sb strings.Builder
			err := component.Render(context.Background(), &sb)
			if err != nil {
				t.Fatalf("failed to render component: %v", err)
			}

			html := sb.String()
			if !strings.Contains(html, tt.message) {
				t.Errorf("expected message %q in output, got: %s", tt.message, html)
			}
			if !strings.Contains(html, `data-level="`+tt.expectedLevel.String()+`"`) {
				t.Errorf("expected level %q in output, got: %s", tt.expectedLevel, html)
			}
		})
	}
}

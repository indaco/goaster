package goaster

import (
	"github.com/a-h/templ"
)

// Toaster struct holds configuration for toast notifications, including styling and behavior options.
type Toaster struct {
	Variant     Variant          // The style variant for the toast notification.
	Border      bool             // Whether to display a border around the toast.
	Rounded     bool             // Whether to display a rounded border for the toast.
	ShowIcon    bool             // Whether to display an icon in the toast.
	Button      bool             // Whether to display a close button.
	AutoDismiss bool             // Whether the toast should automatically dismiss after a delay.
	Animation   bool             // Whether the toast should use animations entrance and exit.
	ProgressBar bool             // Whether the toast should display a progress bar to indicate the remaining time before a toast notification disappears
	Position    Position         // Position of the toast on the screen.
	Icons       map[Level]string // Custom icons for each level, stored as SVG strings.
	queue       *Queue
}

var (
	// cssByLevel maps each Level to a corresponding CSS class for styling.
	mapCssClassByLevel map[Level]templ.CSSClass

	// entranceCssClassesByPosition maps each Position to a corresponding templ.CSSClass representing the entrance direction for toast animations.
	entranceCssClassesByPosition map[Position]templ.CSSClass
)

func init() {
	mapCssClassByLevel = getMapCssClassByLevel()
	entranceCssClassesByPosition = getToastEntranceCSSClassesByPosition()
}

// NewToaster creates a new Toaster instance with default settings, applying any provided options.
func NewToaster(options ...Option) *Toaster {
	toaster := &Toaster{
		Border:      true,
		Rounded:     true,
		ShowIcon:    true,
		Button:      true,
		AutoDismiss: true,
		Animation:   true,
		ProgressBar: true,
		Position:    BottomRight,
		Icons:       defaultIcons(),
		queue:       NewQueue(),
	}

	// Apply each provided configuration option to the Toaster instance.
	for _, option := range options {
		option(toaster)
	}

	return toaster
}

func (t *Toaster) Queue() *Queue {
	return t.queue
}

// RenderAll prints all the toast notifications in the queue.
func (t *Toaster) RenderAll() templ.Component {
	return container(t)
}

// PushDefault adds a default toast notification to the queue.
func (t *Toaster) PushDefault(message string) {
	t.queue.Enqueue(NewToast(message, DefaultLevel))
}

// PushSuccess adds a success toast notification to the queue.
func (t *Toaster) PushSuccess(message string) {
	t.queue.Enqueue(NewToast(message, SuccessLevel))
}

// PushError adds an error toast notification to the queue.
func (t *Toaster) PushError(message string) {
	t.queue.Enqueue(NewToast(message, ErrorLevel))
}

// PushWarning adds a warning toast notification to the queue.
func (t *Toaster) PushWarning(message string) {
	t.queue.Enqueue(NewToast(message, WarningLevel))
}

// PushInfo adds an info toast notification to the queue.
func (t *Toaster) PushInfo(message string) {
	t.queue.Enqueue(NewToast(message, InfoLevel))
}

// Default displays a default toast notification.
func (t *Toaster) Default(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, DefaultLevel))
	return container(t)
}

// Success displays a success toast notification.
func (t *Toaster) Success(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, SuccessLevel))
	return container(t)
}

// Error displays an error toast notification.
func (t *Toaster) Error(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, ErrorLevel))
	return container(t)
}

// Warning displays a warning toast notification.
func (t *Toaster) Warning(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, WarningLevel))
	return container(t)
}

// Info displays an info toast notification.
func (t *Toaster) Info(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, InfoLevel))
	return container(t)
}

package goaster

import "github.com/a-h/templ"

// Toaster holds configuration for toast notifications.
type Toaster struct {
	Variant     Variant          // Style variant
	Border      bool             // Display border
	Rounded     bool             // Use rounded corners
	ShowIcon    bool             // Show icon
	Button      bool             // Show close button
	AutoDismiss bool             // Auto-dismiss after timeout
	Animation   bool             // Animate entrance/exit
	ProgressBar bool             // Show progress bar
	Position    Position         // Screen position
	Icons       map[Level]string // Custom icons
	queue       *Queue           // Toast queue
}

// NewToaster creates a Toaster instance with default settings and applies any given options.
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

	for _, option := range options {
		option(toaster)
	}

	return toaster
}

/* ------------------------------------------------------------------------- */
/* QUEUE ACCESS                                                              */
/* ------------------------------------------------------------------------- */

// Queue returns the internal toast queue.
func (t *Toaster) Queue() *Queue {
	return t.queue
}

/* ------------------------------------------------------------------------- */
/* PUSH METHODS                                                              */
/* ------------------------------------------------------------------------- */

// PushDefault adds a default toast notification.
func (t *Toaster) PushDefault(message string) {
	t.queue.Enqueue(NewToast(message, DefaultLevel))
}

// PushSuccess adds a success toast notification.
func (t *Toaster) PushSuccess(message string) {
	t.queue.Enqueue(NewToast(message, SuccessLevel))
}

// PushError adds an error toast notification.
func (t *Toaster) PushError(message string) {
	t.queue.Enqueue(NewToast(message, ErrorLevel))
}

// PushWarning adds a warning toast notification.
func (t *Toaster) PushWarning(message string) {
	t.queue.Enqueue(NewToast(message, WarningLevel))
}

// PushInfo adds an info toast notification.
func (t *Toaster) PushInfo(message string) {
	t.queue.Enqueue(NewToast(message, InfoLevel))
}

/* ------------------------------------------------------------------------- */
/* RENDER METHODS                                                            */
/* ------------------------------------------------------------------------- */

// RenderAll prints all the toast notifications in the queue.
func (t *Toaster) RenderAll() templ.Component {
	return container(t)
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

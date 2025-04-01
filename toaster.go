package goaster

import (
	"github.com/a-h/templ"
	"github.com/indaco/goaster/components"
	"github.com/indaco/goaster/internal/viewmodel"
)

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

func ToasterDefaults() *Toaster {
	return &Toaster{
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
}

func (t *Toaster) ToViewModel() viewmodel.ToasterViewModel {
	var toasts []viewmodel.ToastViewModel
	for _, toast := range t.Queue().GetMessagesAndDequeue() {
		toasts = append(toasts, viewmodel.ToastViewModel{
			Message: toast.Message,
			Level:   toast.Level.String(),
			Icon:    t.Icons[toast.Level],
		})
	}

	return viewmodel.ToasterViewModel{
		Variant:     t.Variant.String(),
		Border:      t.Border,
		Rounded:     t.Rounded,
		ShowIcon:    t.ShowIcon,
		Button:      t.Button,
		AutoDismiss: t.AutoDismiss,
		Animation:   t.Animation,
		ProgressBar: t.ProgressBar,
		Position:    t.Position.String(),
		Toasts:      toasts,
	}
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
	return components.Container(t.ToViewModel())
}

// Default displays a default toast notification.
func (t *Toaster) Default(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, DefaultLevel))
	return components.Container(t.ToViewModel())
}

// Success displays a success toast notification.
func (t *Toaster) Success(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, SuccessLevel))
	return components.Container(t.ToViewModel())
}

// Error displays an error toast notification.
func (t *Toaster) Error(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, ErrorLevel))
	return components.Container(t.ToViewModel())
}

// Warning displays a warning toast notification.
func (t *Toaster) Warning(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, WarningLevel))
	return components.Container(t.ToViewModel())
}

// Info displays an info toast notification.
func (t *Toaster) Info(message string) templ.Component {
	t.queue.Enqueue(NewToast(message, InfoLevel))
	return components.Container(t.ToViewModel())
}

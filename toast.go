package goaster

type Toast struct {
	Message string
	Level   Level // Assuming ToastLevel is defined as before
}

// NewToast creates and returns a new Message with the specified text and level.
func NewToast(msg string, level Level) Toast {
	return Toast{
		Message: msg,
		Level:   level,
	}
}

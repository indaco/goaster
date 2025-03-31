package goaster

// Toast holds a message and its level.
type Toast struct {
	Message string
	Level   Level
}

// NewToast creates a new Toast with a message and level.
func NewToast(msg string, level Level) Toast {
	return Toast{
		Message: msg,
		Level:   level,
	}
}

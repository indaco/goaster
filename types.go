package goaster

// Level represents the severity level of a toast notification.
type Level string

func (l Level) String() string {
	return string(l)
}

// Position represents the placement for the toast container.
type Position string

func (p Position) String() string {
	return string(p)
}

// Variant represents a style variant for the toast component.
type Variant string

func (v Variant) String() string {
	return string(v)
}

// Icon is a named type for SVG string representations.
type Icon string

func (i Icon) String() string {
	return string(i)
}

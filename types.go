package goaster

import (
	"github.com/a-h/templ"
)

// Level represents the severity level of a toast notification.
type Level string

// Position represents the name and CSS class (templ.CSSClass) for positioning the toast container on the screen.
type Position struct {
	Name string
	CSS  templ.CSSClass
}

// LevelCSSClass represents the CSS class (templ.CSSClass) associated with a specific toast Level.
type LevelCSSClass = templ.CSSClass

// Variant represent a style variant for the toast component.
type Variant string

func (v Variant) String() string {
	return string(v)
}

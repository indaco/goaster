package types

import "html/template"

// PageData represents the data to be rendered in the HTML template
type PageData struct {
	Toast ToastComponent
}

type ToastComponent struct {
	HTML template.HTML
}

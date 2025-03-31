// Package goaster - Helpers for using toasts with `template/html`
package goaster

import (
	"context"
	"fmt"
	"html/template"

	"github.com/a-h/templ"
)

// HTMLGenerator provides functions for generating HTML code for toast notifications.
type HTMLGenerator struct{}

// NewHTMLGenerator creates a new instance of HTMLGenerator.
func NewHTMLGenerator() *HTMLGenerator {
	return &HTMLGenerator{}
}

// DisplayAll generates HTML code for displaying the toast and returns it as a template.HTML.
func (g *HTMLGenerator) DisplayAll(t *Toaster) (template.HTML, error) {
	// Generate HTML code for displaying the toast.
	html, err := templ.ToGoHTML(context.Background(), t.RenderAll())
	if err != nil {
		return "", fmt.Errorf("failed to generate toast HTML: %v", err)
	}
	return html, nil
}

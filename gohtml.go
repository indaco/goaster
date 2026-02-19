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

// DisplayAll generates HTML code for displaying all queued toasts and returns it as a template.HTML.
// The caller must supply a context.Context for the underlying templ rendering.
func (g *HTMLGenerator) DisplayAll(ctx context.Context, t *Toaster) (template.HTML, error) {
	// Generate HTML code for displaying the toast.
	html, err := templ.ToGoHTML(ctx, t.RenderAll())
	if err != nil {
		return "", fmt.Errorf("failed to generate toast HTML: %w", err)
	}
	return html, nil
}

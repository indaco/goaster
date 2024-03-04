// Package goaster - Helpers for using toasts with `template/html`
package goaster

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"html/template"
)

// HTMLGenerator provides functions for generating HTML code for toast notifications.
type HTMLGenerator struct{}

// NewHTMLGenerator creates a new instance of HTMLGenerator.
func NewHTMLGenerator() *HTMLGenerator {
	return &HTMLGenerator{}
}

// GoasterCSSToGoHTML generates HTML code for the toast CSS and returns it as a template.HTML.
func (g *HTMLGenerator) GoasterCSSToGoHTML() (template.HTML, error) {
	html, err := templ.ToGoHTML(context.Background(), GoasterCSS())
	if err != nil {
		return "", fmt.Errorf("failed to generate toast CSS: %v", err)
	}
	return html, nil
}

// GoasterJSToGoHTML generates HTML code for the toast JS and returns it as a template.HTML.
func (g *HTMLGenerator) GoasterJSToGoHTML(t *Toaster, jsOptions *GoasterJSOptions) (template.HTML, error) {
	html, err := templ.ToGoHTML(context.Background(), GoasterJS(t, jsOptions))
	if err != nil {
		return "", fmt.Errorf("failed to generate toast CSS: %v", err)
	}
	return html, nil
}

// DisplayAll generates HTML code for displaying the toast and returns it as a template.HTML.
func (g *HTMLGenerator) DisplayAll(t *Toaster) (template.HTML, error) {
	// Generate HTML code for displaying the toast.
	html, err := templ.ToGoHTML(context.Background(), t.DisplayAll())
	if err != nil {
		return "", fmt.Errorf("failed to generate toast JS: %v", err)
	}
	return html, nil
}

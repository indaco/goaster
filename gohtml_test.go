package goaster

import (
	"html/template"
	"testing"
)

func TestGoasterCSSToGoHTML(t *testing.T) {
	g := NewHTMLGenerator()

	html, err := g.GoasterCSSToGoHTML()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if html == "" {
		t.Error("expected non-empty HTML output for CSS")
	}
}

func TestGoasterJSToGoHTML(t *testing.T) {
	g := NewHTMLGenerator()

	toaster := NewToaster()
	jsHTML, err := g.GoasterJSToGoHTML(toaster, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if jsHTML == "" {
		t.Error("expected non-empty HTML output for JS")
	}
}

func TestDisplayAll(t *testing.T) {
	g := NewHTMLGenerator()

	toaster := NewToaster()
	toaster.PushInfo("Sample Toast")

	html, err := g.DisplayAll(toaster)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if html == "" {
		t.Error("expected non-empty HTML output for DisplayAll")
	}
}

func TestGeneratedHTML_IsTemplateHTML(t *testing.T) {
	g := NewHTMLGenerator()

	toaster := NewToaster()
	toaster.PushSuccess("Success Toast")

	html, err := g.DisplayAll(toaster)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that it's the correct type
	_, ok := interface{}(html).(template.HTML)
	if !ok {
		t.Errorf("expected template.HTML, got %T", html)
	}
}

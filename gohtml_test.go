package goaster

import (
	"html/template"
	"testing"
)

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
	_, ok := any(html).(template.HTML)
	if !ok {
		t.Errorf("expected template.HTML, got %T", html)
	}
}

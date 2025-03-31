package handlers

import (
	"net/http"

	"github.com/indaco/goaster/examples/pages"

	"github.com/a-h/templ"
	"github.com/indaco/goaster"
)

var toaster *goaster.Toaster

func init() {
	// Create a toaster instance
	toaster = goaster.NewToaster(goaster.WithAutoDismiss(false))
	// Push messages
	toaster.PushDefault("Default Toast")
	toaster.PushSuccess("Success Toast")
	toaster.PushError("Error Toast")
	toaster.PushWarning("Warning Toast")
	toaster.PushInfo("Info Toast")
}

func render(w http.ResponseWriter, r *http.Request, statusCode int, t templ.Component) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/html")
	return t.Render(r.Context(), w)
}

func HandleMultiple(w http.ResponseWriter, r *http.Request) {
	err := render(w, r, http.StatusOK, pages.MultipleToastPage(toaster))
	if err != nil {
		return
	}
}

package handlers

import (
	"net/http"

	"github.com/indaco/goaster/examples/pages"

	"github.com/a-h/templ"
	"github.com/indaco/goaster"
)

func HandleVariantNeon(w http.ResponseWriter, r *http.Request) {
	toaster := goaster.NewToaster(
		goaster.WithVariant(goaster.Neon),
		goaster.WithAutoDismiss(false),
	)
	toaster.PushDefault("Default Toast")
	toaster.PushSuccess("Success Toast")
	toaster.PushError("Error Toast")
	toaster.PushWarning("Warning Toast")
	toaster.PushInfo("Info Toast")
	templ.Handler(pages.VariantNeonPage(toaster)).ServeHTTP(w, r)
}

package handlers

import (
	"net/http"

	"github.com/indaco/goaster/demos/pages"

	"github.com/a-h/templ"
	"github.com/indaco/goaster"
)

func HandleVariantAccentLight(w http.ResponseWriter, r *http.Request) {
	toaster := goaster.NewToaster(
		goaster.WithVariant(goaster.AccentLight),
		goaster.WithAutoDismiss(false),
	)
	templ.Handler(pages.VariantAccentLightPage(toaster)).ServeHTTP(w, r)
}

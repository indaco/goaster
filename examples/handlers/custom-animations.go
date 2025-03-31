package handlers

import (
	"net/http"

	"github.com/indaco/goaster/examples/pages"

	"github.com/a-h/templ"
	"github.com/indaco/goaster"
)

func HandleCustomAnimations(w http.ResponseWriter, r *http.Request) {
	toaster := goaster.NewToaster(
		goaster.WithAutoDismiss(false),
	)
	templ.Handler(pages.CustomAnimationPage(toaster)).ServeHTTP(w, r)
}

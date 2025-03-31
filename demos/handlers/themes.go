package handlers

import (
	"net/http"

	"github.com/indaco/goaster/demos/pages"

	"github.com/a-h/templ"
	"github.com/indaco/goaster"
)

func HandleThemes(w http.ResponseWriter, r *http.Request) {
	toaster := goaster.NewToaster()
	templ.Handler(pages.ThemesPage(toaster)).ServeHTTP(w, r)
}

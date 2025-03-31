package handlers

import (
	"net/http"

	"github.com/indaco/goaster/demos/pages"

	"github.com/a-h/templ"
	"github.com/indaco/goaster"
)

func HandleSingle(w http.ResponseWriter, r *http.Request) {
	result := true
	toaster := goaster.NewToaster()
	templ.Handler(pages.SingleToastPage(result, toaster)).ServeHTTP(w, r)
}

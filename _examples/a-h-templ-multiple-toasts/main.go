package main

import (
	"log"
	"net/http"
	"os"

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

func HandleHome(w http.ResponseWriter, r *http.Request) {
	err := render(w, r, http.StatusOK, HomePage(toaster))
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", HandleHome)

	port := ":3300"
	log.Printf("Listening on %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
}

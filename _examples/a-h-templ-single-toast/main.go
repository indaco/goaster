package main

import (
	"github.com/a-h/templ"
	"github.com/indaco/goaster"
	"log"
	"net/http"
	"os"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	result := true
	toaster := goaster.NewToaster()
	templ.Handler(HomePage(result, toaster)).ServeHTTP(w, r)
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

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/indaco/goaster/examples/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.HandleHome)
	mux.HandleFunc("GET /single", handlers.HandleSingle)
	mux.HandleFunc("GET /multiple", handlers.HandleMultiple)
	mux.HandleFunc("GET /custom-animations", handlers.HandleCustomAnimations)
	mux.HandleFunc("GET /custom-icons", handlers.HandleCustomIcons)
	mux.HandleFunc("GET /themes", handlers.HandleThemes)
	mux.HandleFunc("GET /variants/accent-light", handlers.HandleVariantAccentLight)
	mux.HandleFunc("GET /gohtml/single", handlers.HandleGoHtmlSingle)
	mux.HandleFunc("GET /gohtml/multiple", handlers.HandleGoHtmlMultiple)

	port := ":3300"
	log.Printf("Listening on %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
}

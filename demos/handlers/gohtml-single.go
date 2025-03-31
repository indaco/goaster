package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/indaco/goaster/demos/types"

	"github.com/a-h/templ"
	"github.com/indaco/goaster"
)

func HandleGoHtmlSingle(w http.ResponseWriter, r *http.Request) {
	// Create a toaster instance
	toaster := goaster.NewToaster()

	toastHTMLGenerator := goaster.NewHTMLGenerator()

	// Render the needed CSS for toast component as template.HTML
	toastCSS, err := toastHTMLGenerator.GoasterCSSToGoHTML()
	if err != nil {
		http.Error(w, "failed to render toast CSS", http.StatusInternalServerError)
		log.Printf("failed to render toast CSS: %v", err)
		return
	}

	// Render the toast component into a template.HTML
	c := toaster.Success("Success Toast")

	toastHtml, err := templ.ToGoHTML(context.Background(), c)
	if err != nil {
		http.Error(w, "failed to generate toast HTML", http.StatusInternalServerError)
		log.Printf("failed to generate toast HTML: %v", err)
		return
	}

	// Render the needed JS for toast component as template.HTML
	toastJS, err := toastHTMLGenerator.GoasterJSToGoHTML(toaster, nil)
	if err != nil {
		http.Error(w, "failed to render toast JS", http.StatusInternalServerError)
		log.Printf("failed to render toast JS: %v", err)
		return
	}

	data := types.PageData{
		Toast: types.ToastComponent{
			CSS:  toastCSS,
			HTML: toastHtml,
			JS:   toastJS,
		},
	}

	// Parse the HTML template
	tmpl := template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>goaster - template/html 1</title>
			<!-- inject toast styles -->
			{{ .Toast.CSS }}
		</head>
		<body>
			<!-- display the messages -->
			{{ .Toast.HTML }}

			<!-- inject toast javascript -->
			{{ .Toast.JS }}
		</body>
		</html>
	`))

	// Execute the template with the provided data and write the output to the response writer
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		log.Printf("failed to render template: %v", err)
		return
	}
}

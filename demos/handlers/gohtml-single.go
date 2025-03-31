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

	// Render the toast component into a template.HTML
	c := toaster.Success("Success Toast")

	toastHtml, err := templ.ToGoHTML(context.Background(), c)
	if err != nil {
		http.Error(w, "failed to generate toast HTML", http.StatusInternalServerError)
		log.Printf("failed to generate toast HTML: %v", err)
		return
	}

	data := types.PageData{
		Toast: types.ToastComponent{
			HTML: toastHtml,
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
		</head>
		<body>
			<!-- display the messages -->
			{{ .Toast.HTML }}
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

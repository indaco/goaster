package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/indaco/goaster/examples/types"

	"github.com/indaco/goaster"
)

func HandleGoHtmlMultiple(w http.ResponseWriter, r *http.Request) {

	// Create a toaster instance
	toaster := goaster.NewToaster(goaster.WithAutoDismiss(false))

	toastHTMLGenerator := goaster.NewHTMLGenerator()

	// Push messages
	toaster.PushDefault("Default Toast")
	toaster.PushSuccess("Success Toast")
	toaster.PushError("Error Toast")
	toaster.PushWarning("Warning Toast")
	toaster.PushInfo("Info Toast")

	// Render the toast component into a template.HTML
	toastHtml, err := toastHTMLGenerator.DisplayAll(toaster)
	if err != nil {
		http.Error(w, "failed to render toast HTML", http.StatusInternalServerError)
		log.Printf("failed to render toast HTML: %v", err)
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
			<title>go-templ-toast</title>
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

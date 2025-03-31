package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/indaco/goaster/demos/types"

	"github.com/indaco/goaster"
)

func HandleGoHtmlMultiple(w http.ResponseWriter, r *http.Request) {

	// Create a toaster instance
	toaster := goaster.NewToaster(goaster.WithAutoDismiss(false))

	toastHTMLGenerator := goaster.NewHTMLGenerator()

	// Render the needed CSS for toast component as template.HTML
	toastCSS, _ := toastHTMLGenerator.GoasterCSSToGoHTML()

	// Push messages
	toaster.PushDefault("Default Toast")
	toaster.PushSuccess("Success Toast")
	toaster.PushError("Error Toast")
	toaster.PushWarning("Warning Toast")
	toaster.PushInfo("Info Toast")

	// Render the toast component into a template.HTML
	toastHtml, err := toastHTMLGenerator.DisplayAll(toaster)

	// Render the needed JS for toast component as template.HTML
	toastJS, _ := toastHTMLGenerator.GoasterJSToGoHTML(toaster, nil)

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
			<title>go-templ-toast</title>
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

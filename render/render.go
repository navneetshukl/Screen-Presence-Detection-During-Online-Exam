package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, url string) {
	tmpl, _ := template.ParseFiles(url)

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Error in parsing template")
	}
}
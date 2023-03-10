package main

import (
	"go_modules/routes"
	"net/http"

	"github.com/go-chi/chi"
)



func main() {
	mux := chi.NewRouter()
	mux.Get("/",routes.Home)
	mux.Post("/",routes.GetData)
	mux.Get("/save",routes.WebCam)
	mux.Post("/save",routes.Save) 
	mux.Get("/display",routes.Display)
	mux.Get("/getData",routes.GetAllData)
	http.ListenAndServe(":8000",mux)
}
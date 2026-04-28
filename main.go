package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/darekvu/lenslocked/controllers"
	"github.com/darekvu/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))
	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))
	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})
	fmt.Println("Starting Server on port 3000...")
	http.ListenAndServe("127.0.0.1:3000", r)
}

package record

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Routes set route
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{txid}", emptyservice)
	router.Post("/", emptyservice)
	return router
}

func emptyservice(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "test")
}

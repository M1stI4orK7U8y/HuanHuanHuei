package record

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/service"
)

// Routes set route
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", getRecord)
	router.Post("/", getRecords)
	return router
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(strings.TrimSpace(id)) == 0 {
		http.Error(w, http.StatusText(400)+": no id", 400)
	}
	ret, err := service.GetRecord(id)
	if err != nil {
		http.Error(w, http.StatusText(500)+": "+err.Error(), 500)
	} else {
		render.JSON(w, r, ret)
	}
}

func getRecords(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ret, err := service.GetRecords(r.Form["ids"])
	if err != nil {
		http.Error(w, http.StatusText(500)+": "+err.Error(), 500)
	} else {
		render.JSON(w, r, ret)
	}

}

package huanhuan

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"gitlab.com/packtumi9722/huanhuanhuei/src/database/model/token"
	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/service"
)

// Routes set route
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", NewHuanHuanJob)
	return router
}

// NewHuanHuanJob request a new huanhuan job
func NewHuanHuanJob(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//NewHuanHuanJob(txid, receiver string, from, to token.TokenType)
	txid := r.FormValue("txid")
	receiver := r.FormValue("receiver")
	if len(txid) == 0 || len(receiver) == 0 {
		http.Error(w, http.StatusText(500)+": txid or receiver not provided", 500)
		return
	}

	from, parseerr := strconv.Atoi(r.FormValue("from"))
	if parseerr != nil {
		http.Error(w, http.StatusText(500)+": "+parseerr.Error(), 500)
		return
	}

	to, parseerr := strconv.Atoi(r.FormValue("to"))
	if parseerr != nil {
		http.Error(w, http.StatusText(500)+": "+parseerr.Error(), 500)
		return
	}

	ret, err := service.NewHuanHuanJob(txid, receiver, token.TokenType(from), token.TokenType(to))
	if err != nil {
		http.Error(w, http.StatusText(500)+": "+err.Error(), 500)
	} else {
		render.JSON(w, r, ret)
	}
}

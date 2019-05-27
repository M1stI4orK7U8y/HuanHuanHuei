package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/api/http/record"
)

// NewRoute get new route
func NewRoute() *chi.Mux {
	return routes()
}

// Routes : initial Routes
func routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		//render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/record", record.Routes())
	})

	return router
}

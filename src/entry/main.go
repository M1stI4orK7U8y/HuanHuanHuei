package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	r "gitlab.com/packtumi9722/huanhuanhuei/src/entry/api/http"
	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/config"
)

func main() {
	router := r.NewRoute()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}
	log.Fatal(http.ListenAndServe(":"+config.Port(), router)) // Note, the port is usually gotten from the environment.
}

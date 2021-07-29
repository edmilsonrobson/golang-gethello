package main

import (
	"net/http"

	"github.com/edmilsonrobson/golang-gethello/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/hello", handlers.SayHello())

	return mux
}

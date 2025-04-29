package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Body struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler() http.Handler {
	route := chi.NewMux()

	route.Use(middleware.Recoverer)
	route.Use(middleware.RequestID)
	route.Use(middleware.Logger)

	route.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	route.Group(func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/create", CreateLink)
			r.Get("/{code}", GetLink)
		})
	})

	return route
}

func CreateLink(w http.ResponseWriter, r *http.Request) {

}

func GetLink(w http.ResponseWriter, r *http.Request) {

}

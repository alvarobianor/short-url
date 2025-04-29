package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"

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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewHandler(db map[string]string) http.Handler {
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
			r.Get("/create", CreateLink(db))
			r.Get("/{code}", GetLink(db))
		})
	})

	return route
}

func CreateLink(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body Body

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			SendJson(w, Response{Error: err.Error()}, http.StatusUnprocessableEntity)
			return
		}

		_, errorUrl := url.Parse(body.URL)

		if errorUrl != nil {
			SendJson(w, Response{Error: errorUrl.Error()}, http.StatusBadRequest)
			return
		}

		code := fmt.Sprintf("%x", md5.Sum([]byte(body.URL)))[:6]

		db[code] = body.URL

		SendJson(w, Response{Data: map[string]string{"code": code}}, http.StatusCreated)

		code = generateCode()

		db[code] = body.URL

		SendJson(w, Response{Data: code}, http.StatusCreated)
	}
}

func generateCode() string {
	lenghtBytes := 16
	byts := make([]byte, lenghtBytes)

	for i := range byts {
		byts[i] = charset[rand.IntN(len(charset))]
	}

	return string(byts)
}

func GetLink(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func SendJson(w http.ResponseWriter, r Response, code int) {
	data, err := json.Marshal(r)

	if err != nil {
		slog.Error(err.Error(), "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)

	if err != nil {
		slog.Error(err.Error(), "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

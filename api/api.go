package api

import (
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
const version = "v1"

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
		r.Route(fmt.Sprintf("/%s", version), func(r chi.Router) {
			r.Post("/create", CreateLink(db))
			r.Get("/{code}", GetLink(db))
		})
	})

	return route
}

func CreateLink(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body Body

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			slog.Error(err.Error(), "error", err)
			SendJson(w, Response{Error: err.Error()}, http.StatusUnprocessableEntity)
			return
		}

		_, errorUrl := url.Parse(body.URL)

		if errorUrl != nil {
			SendJson(w, Response{Error: errorUrl.Error()}, http.StatusBadRequest)
			return
		}

		code := generateCode(db)

		db[code] = body.URL

		url := fmt.Sprintf("%s/%s/%s", r.Host, version, code)

		SendJson(w, Response{Data: url}, http.StatusCreated)
	}
}

func generateCode(db map[string]string) string {
	lenghtBytes := 16
	byts := make([]byte, lenghtBytes)

	for i := range byts {
		byts[i] = charset[rand.IntN(len(charset))]
	}

	if _, ok := db[string(byts)]; ok {
		return generateCode(db)
	}

	return string(byts)
}

func GetLink(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		if code == "" {
			SendJson(w, Response{Error: "code is required"}, http.StatusBadRequest)
			return
		}

		url, ok := db[code]

		if !ok {
			SendJson(w, Response{Error: "code not found"}, http.StatusNotFound)
			return
		}

		SendJson(w, Response{Data: url}, http.StatusOK)

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

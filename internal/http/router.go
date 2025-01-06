package http

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *Handler) http.Handler {
	r := chi.NewRouter()
	r.Post("/shorten", handler.CreateShortURL)
	r.Get("/{short_url}", handler.Redirect)
	return r
}
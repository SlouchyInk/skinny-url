package http

import (
	"encoding/json"
	"net/http"

	"github.com/hdurham99/skinny-url/internal/shortener"
)

type Handler struct {
	Service *shortener.Service
}

func NewHandler(service *shortener.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateShortURL(writer http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	code, err := h.Service.ShortenURL(req.URL)
	if err != nil {
		http.Error(writer, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(map[string]string{"short_url": code})
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	url, err := h.Service.GetOriginalURL(code)

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

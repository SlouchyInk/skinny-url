package http

import (
	"encoding/json"
	"fmt"
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
		URL     string `json:"url"`
		USER_ID string `json:"user_id"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	code, err := h.Service.ShortenURL(req.URL, req.USER_ID)
	if err != nil {
		http.Error(writer, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(map[string]string{"short_url": code})
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	url, err := h.Service.GetOriginalURL(code)
	fmt.Println(code)
	fmt.Println(url)

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	h.Service.IncrementClickCount(code)
	http.Redirect(w, r, url, http.StatusFound)
}

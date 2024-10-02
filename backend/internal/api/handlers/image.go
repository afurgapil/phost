package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/afurgapil/phost/backend/pkg/entities"
	"github.com/afurgapil/phost/backend/pkg/image"
)

type Handler struct {
	service image.Service
}

func NewHandler(service image.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateImage(w, r)
	case http.MethodGet:
		h.GetImageByID(w, r)
	case http.MethodDelete:
		h.DeleteImage(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CreateImage(w http.ResponseWriter, r *http.Request) {
	var image entities.Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	image.ID = 0

	createdImage, err := h.service.CreateImage(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdImage)
}

func (h *Handler) GetImageByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing image ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}

	image, err := h.service.GetImageByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(image)
}

func (h *Handler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing image ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteImage(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

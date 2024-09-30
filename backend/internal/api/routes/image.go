package routes

import (
	"net/http"

	"github.com/afurgapil/phost/backend/internal/api/handlers"
)

func RegisterRoutes(handler *handlers.Handler) {
	http.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateImage(w, r)
		case http.MethodGet:
			handler.GetImageByID(w, r)
		case http.MethodDelete:
			handler.DeleteImage(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

package main

import (
	"log"
	"net/http"

	"github.com/afurgapil/phost/backend/internal/api/handlers"
	"github.com/afurgapil/phost/backend/internal/api/routes"
	"github.com/afurgapil/phost/backend/internal/config"
	"github.com/afurgapil/phost/backend/pkg/image"
)

func main() {
	port, loadErr := config.LoadConfig("PORT")
	if loadErr != nil || port == "" {
		port = "8081"
	}

	dbPort, loadErr := config.LoadConfig("DB_PORT")
	if loadErr != nil || dbPort == "" {
		dbPort = "8080"
	}

	imageRepository := image.NewRepository("http://localhost:" + dbPort)
	imageService := image.NewService(imageRepository)
	handler := handlers.NewHandler(imageService)
	routes.RegisterRoutes(handler)

	log.Printf("Starting backend server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

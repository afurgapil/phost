package main

import (
	"log"
	"net/http"

	"github.com/afurgapil/phost/backend/internal/api/handlers"
	"github.com/afurgapil/phost/backend/internal/config"
	"github.com/afurgapil/phost/backend/pkg/image"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(handler)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        corsHandler,
		MaxHeaderBytes: 1 << 30,
	}

	log.Printf("Starting backend server on port %s...", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

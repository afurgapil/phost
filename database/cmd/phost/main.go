package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/afurgapil/phost/database/internal/config"
	"github.com/afurgapil/phost/database/internal/database"
	"github.com/afurgapil/phost/database/pkg/handler"
)

func main() {
	db := &database.Database{}

	port, loadErr := config.LoadConfig("PORT")
	if port == "" || loadErr != nil {
		port = "8080"
	}

	err := db.LoadFromFile("data.json")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Data loading error: %v", err)
		}
		log.Println("File not found, a new file will be created.")
	}

	handler.SetDatabase(db)

	http.HandleFunc("/execute", handler.HandleExecute)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting server on port %s...", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
		}
	}()

	<-signalChan
	log.Println("The server has received a shutdown signal, we are logging data...")

	err = db.SaveToFile("data.json")
	if err != nil {
		log.Fatalf("Data save error: %v", err)
	}

	log.Println("The server was shut down and the data was saved.")
}

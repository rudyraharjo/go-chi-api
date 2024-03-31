package main

import (
	"context"
	"log"

	"github.com/rudyraharjo/api-rud/application"
	"github.com/rudyraharjo/api-rud/database"
)

func main() {

	configFile := "config-development.yaml"
	if err := application.LoadConfig(configFile); err != nil {
		panic(err)
	}

	cfg := application.Cfg

	// Connect to PostgreSQL
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		panic(err)
	}

	app, err := application.NewApp(cfg, db)
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	// Jalankan aplikasi (mulai server HTTP)
	if err := app.Serve(context.Background()); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

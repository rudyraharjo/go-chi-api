package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	config Config
	db     *pgxpool.Pool
}

func NewApp(config Config, db *pgxpool.Pool) (*App, error) {

	app := &App{
		db:     db,
		config: config,
	}

	return app, nil
}

func (app *App) Serve(ctx context.Context) error {
	port := app.config.App.Port

	log.Println("API listening on port", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.Routes(app.db),
	}

	conn, err := app.db.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	defer conn.Release()

	// Uji koneksi database
	if err := conn.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	fmt.Println("Starting server")

	return server.ListenAndServe()

}

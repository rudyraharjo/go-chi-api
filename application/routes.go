package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (a *App) Routes(db *pgxpool.Pool) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// all v1 routes
	router.Route("/api/v1/", func(router chi.Router) {
		router.Get("/auth/login", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("auth/login!"))
		})
		router.Get("/auth/signup", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("auth/signup!"))
		})
	})

	return router
}

package main

import (
	project "gothstarter"
	"gothstarter/internal/config"
	"gothstarter/internal/handlers"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.NotFound(handlers.Make(handlers.HandleNotFound))

	router.Handle("/static/*", http.StripPrefix("/static", project.Public()))

	router.Get("/", handlers.Make(handlers.HandleHome))

	cfg := config.LoadConfig()
	slog.Info("HTTP server started", "listenAddr", cfg.Port)
	http.ListenAndServe(cfg.Port, router)

}

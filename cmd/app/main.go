package main

import (
	"fmt"
	project "gothstarter"
	"gothstarter/internal/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Handle("/static/*", project.Public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	fmt.Println("hello world!")

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)

}

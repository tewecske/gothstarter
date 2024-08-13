package main

import (
	"context"
	"errors"
	project "gothstarter"
	"gothstarter/internal/config"
	"gothstarter/internal/handlers"
	"gothstarter/internal/hash/passwordhash"
	database "gothstarter/internal/store/db"
	"gothstarter/internal/store/session"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := config.LoadConfig()

	passwordHash := passwordhash.NewHPasswordHash()

	dbAccess := database.SetupDB(cfg.DatabaseName, passwordHash)

	dbAccess.UserStore.CreateUser("a@a.a", "aaa")
	dbAccess.SessionStore.CreateSession(&session.Session{
		UserID: 1,
	})

	router := chi.NewRouter()

	router.NotFound(handlers.Make(handlers.HandleNotFound))

	router.Handle("/static/*", http.StripPrefix("/static", project.Public()))

	router.Get("/", handlers.Make(handlers.HandleHome))

	router.Get("/login", handlers.Make(handlers.HandleLogin))

	router.Post("/login", handlers.Make(handlers.HandlePostLogin(
		dbAccess.UserStore,
		dbAccess.SessionStore,
		passwordHash,
		cfg.SessionCookieName,
	)))

	// slog.Info("HTTP server started", "listenAddr", cfg.Port)
	// http.ListenAndServe(cfg.Port, router)

	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Server shutdown complete")
		} else if err != nil {
			logger.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", cfg.Port))
	<-killSig

	logger.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbCloseError := dbAccess.DB.Close()
	if dbCloseError != nil {
		logger.Error("DB close failed", slog.Any("err", dbCloseError))
	} else {
		logger.Info("DB closed")
	}

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Server shutdown complete")
}

package main

import (
	"log/slog"
	"net/http"
	"os"
	"shortUrl/api"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("All systems are offline")
}

func run() error {

	db := make(map[string]string)

	handler := api.NewHandler(db)

	server := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	slog.Info("Starting server ->", "addr", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

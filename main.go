package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	mux := http.NewServeMux()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}
	mux.HandleFunc("/", app.home)

	logMiddleware := app.loggingMiddleware(mux)

	logger.Info("Starting server", "addr", ":4000")
	err := http.ListenAndServe(":4000", logMiddleware)
	logger.Error(err.Error())
	os.Exit(1)
}

package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}


func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let's explore Dependency Injection in Go "))
}


// Middleware function that wraps an http.Handler
func (app *application)loggingMiddleware(next http.Handler) http.Handler {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//gather the log values before creating the logger
		//th Printf() function 
		//var(
		//	ip := r.RemoteAddr
		//	proto := r.Proto
		//	method := r.Method
		//	uri := r.URL.RequestURI()
		//)
		//pre-processing
		app.logger.Info("Pre-processing")
		// Call the next handler
		next.ServeHTTP(w, r)
		// Post-processing log
		app.logger.Info("Post-processing")
	})
	return fn
}

func main() {
	mux := http.NewServeMux()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	//create a new instance of the application struct
	app := &application{
		logger: logger,
	}
	mux.HandleFunc("/", home)

	// Correctly wrapping mux with the middleware
	logMiddleware := app.loggingMiddleware(mux)

	logger.Info("Starting server", "addr", ":4000")
	err := http.ListenAndServe(":4000", logMiddleware)
	logger.Error(err.Error())
	os.Exit(1)
}

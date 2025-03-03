package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let's explore Dependency Injection in Go"))
}

// Middleware function that wraps an http.Handler
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details
		ip := r.RemoteAddr
		proto := r.Proto
		method := r.Method
		uri := r.URL.RequestURI()
		log.Printf("Received request - IP: %s, Protocol: %s, Method: %s, URI: %s", ip, proto, method, uri)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Post-processing log
		log.Println("Request processed")
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Correctly wrapping mux with the middleware
	logMiddleware := loggingMiddleware(mux)

	log.Print("Starting the server on port: 4000")
	err := http.ListenAndServe(":4000", logMiddleware)
	log.Fatal(err)
}

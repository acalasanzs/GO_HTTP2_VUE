package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Create a new Gorilla mux router.
	r := mux.NewRouter()

	// API endpoint that returns a JSON list of todos.
	r.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Sample todo items.
		w.Write([]byte(`[{"id":1,"task":"Buy groceries","completed":false},
		                  {"id":2,"task":"Walk the dog","completed":true},
		                  {"id":3,"task":"Read a book","completed":false}]`))
	}).Methods("GET")

	// Serve static files (the Vue app) from the "public" directory.
	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/").Handler(fs)

	// Configure CORS to allow all origins and common methods.
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	// Wrap the router with the CORS handler.
	handler := c.Handler(r)

	// Start the HTTPS server on port 8443.
	// HTTP/2 is automatically enabled when using TLS.
	log.Println("Starting server on https://localhost:8443")
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", handler)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

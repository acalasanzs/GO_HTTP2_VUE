package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var (
	items     = []Item{}
	itemsLock sync.RWMutex
	nameGen   *NameGenerator
)

type NameGenerator struct {
	adjectives []string
	nouns      []string
	rng        *rand.Rand
}

// Item represents a simple data item
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Generate randomly selects an adjective and noun and combines them.
func (ng *NameGenerator) Generate() string {
	adjective := ng.adjectives[ng.rng.Intn(len(ng.adjectives))]
	noun := ng.nouns[ng.rng.Intn(len(ng.nouns))]
	return adjective + " " + noun
}
func main() {
	godotenv.Load()
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/items", getItems).Methods("GET")
	r.HandleFunc("/api/items", createItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", getItem).Methods("GET")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Vue.js dev server address
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Start server
	handler := c.Handler(r)
	log.Println("Go server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Item not found"})
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

package main

import (
	"climbpass-matching-service/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRouter specified all routes
func InitRouter() *mux.Router {
	r := mux.NewRouter()

	// Serve static files in assets
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	// Routes
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")
	controllers.AddGymControllers(r)

	return r
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

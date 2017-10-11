package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest-in-go/api"
	"fmt"
)

var routes = mux.NewRouter()

// Route Configuration
func initRoutes() {
	// Parent Routes
	routes.HandleFunc("/", api.IndexHandler).Methods("GET", "POST", "PATCH", "OPTIONS", "DELETE", "PUT")

	// Sub-routes { /api/v1 }
	v1 := routes.NewRoute().PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/user/create", api.CreateUser).Methods("POST")

	fmt.Println("Application Running On :8009...")
	http.ListenAndServe(":8009", routes)
}

// Application Start Point
func main() {
	initRoutes()
}

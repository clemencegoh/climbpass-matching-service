package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// AddGymControllers handles routing for gyms endpoints
func AddGymControllers(r *mux.Router) {
	baseRoute := "/gyms"
	r.HandleFunc(baseRoute, getGyms).Methods("GET")
}


// @GetController( path = "/" )
func getGyms(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "From gyms")
}

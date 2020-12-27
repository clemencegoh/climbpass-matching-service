package controllers

import (
	"climbpass-matching-service/repositories"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddGymControllers handles routing for gyms endpoints
func AddGymControllers(r *mux.Router) {
	baseRoute := "/gyms"
	r.HandleFunc(baseRoute, getGyms).Methods("GET")
}

// @GET("/")
func getGyms(w http.ResponseWriter, r *http.Request) {

	repo := repositories.GymRepository{}
	gym := repo.GetGymByName("init_name")
	response, err := json.Marshal(gym)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

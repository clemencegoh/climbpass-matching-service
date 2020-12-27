package controllers

import (
	"climbpass-matching-service/models"
	"climbpass-matching-service/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddGymControllers handles routing for gyms endpoints
func AddGymControllers(r *mux.Router) {

	controller := NewGymController()

	baseRoute := "/gyms"
	r.HandleFunc(baseRoute, controller.getAllGyms).Methods("GET")
	r.HandleFunc(baseRoute+"/{name}", controller.getGymByName).Methods("GET")
	r.HandleFunc(baseRoute, controller.createGym).Methods("POST")
}

type IGymController interface {
	getAllGyms(w http.ResponseWriter, r *http.Request)
	createGym(w http.ResponseWriter, r *http.Request)
	getGymByName(w http.ResponseWriter, r *http.Request)
}

type GymController struct {
	service services.IGymService
}

func NewGymController() IGymController {
	service := services.NewGymService()
	return GymController{service}
}

// @GET("/")
func (controller GymController) getAllGyms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := controller.service.GetAllGyms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

// @GET("/:name")
func (controller GymController) getGymByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["name"]
	response, err := controller.service.GetGymByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

// @POST("/")
func (controller GymController) createGym(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var g models.GymModel
	err := decoder.Decode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := controller.service.CreateGym(g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}

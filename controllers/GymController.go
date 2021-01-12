package controllers

import (
	"climbpass-matching-service/constants"
	"climbpass-matching-service/models"
	"climbpass-matching-service/services"
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

// AddGymControllers handles routing for gyms endpoints
func AddGymControllers(r *mux.Router) {
	controller := NewGymController()
	baseRoute := constants.APIBasePath + "/gyms"

	r.HandleFunc(baseRoute, controller.getAllGyms).Methods("GET")
	r.HandleFunc(baseRoute+"/{name}", controller.getGymByName).Methods("GET")
	r.HandleFunc(baseRoute, controller.createGym).Methods("POST")
	r.HandleFunc(baseRoute+"/{id}", controller.deleteGymByID).Methods("DELETE")
	r.HandleFunc(baseRoute+"/{id}", controller.updateGymByID).Methods("PUT")
}

// IGymController interface for GymController
type IGymController interface {
	getAllGyms(w http.ResponseWriter, r *http.Request)
	createGym(w http.ResponseWriter, r *http.Request)
	getGymByName(w http.ResponseWriter, r *http.Request)
	deleteGymByID(w http.ResponseWriter, r *http.Request)
	updateGymByID(w http.ResponseWriter, r *http.Request)
}

// GymController implements interface
type GymController struct {
	service services.IGymService
}

// NewGymController inits GymController
func NewGymController() IGymController {
	service := services.NewGymService()
	return GymController{service}
}

// @GET("/")
func (controller GymController) getAllGyms(w http.ResponseWriter, r *http.Request) {
	response, err := controller.service.GetAllGyms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

// @GET("/:name")
func (controller GymController) getGymByName(w http.ResponseWriter, r *http.Request) {
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
	decoder := json.NewDecoder(r.Body)
	var g models.GymProfile
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

// @DELETE("/:id")
func (controller GymController) deleteGymByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := controller.service.DeleteGymByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

// @PUT("/:id")
func (controller GymController) updateGymByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var g models.GymProfile
	err2 := decoder.Decode(&g)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	response, err3 := controller.service.UpdateGymByID(id, g)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

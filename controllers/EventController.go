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

// AddEventControllers handles routing for Event endpoints
func AddEventControllers(r *mux.Router) {
	controller := NewEventController()
	baseRoute := constants.APIBasePath + "/events"

	r.HandleFunc(baseRoute, controller.getAllEvent).Methods("GET")
	r.HandleFunc(baseRoute+"/{id}", controller.getEventByID).Methods("GET")
	r.HandleFunc(baseRoute, controller.createEvent).Methods("POST")
	r.HandleFunc(baseRoute+"/{id}", controller.deleteEventByID).Methods("DELETE")
	r.HandleFunc(baseRoute+"/{id}", controller.updateEventByID).Methods("PUT")
}

// IEventController interface for EventController
type IEventController interface {
	getAllEvent(w http.ResponseWriter, r *http.Request)
	createEvent(w http.ResponseWriter, r *http.Request)
	getEventByID(w http.ResponseWriter, r *http.Request)
	deleteEventByID(w http.ResponseWriter, r *http.Request)
	updateEventByID(w http.ResponseWriter, r *http.Request)
}

// EventController implements interface
type EventController struct {
	service services.IEventService
}

// NewEventController inits EventController
func NewEventController() IEventController {
	service := services.NewEventService()
	return EventController{service}
}

// @GET("/")
func (controller EventController) getAllEvent(w http.ResponseWriter, r *http.Request) {
	response, err := controller.service.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

// @GET("/:id")
func (controller EventController) getEventByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err0 := strconv.ParseUint(params["id"], 10, 32)
	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
		return
	}
	response, err := controller.service.GetEventByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

// @POST("/")
func (controller EventController) createEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var m models.EventModel
	err := decoder.Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := controller.service.CreateEvent(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}

// @DELETE("/:id")
func (controller EventController) deleteEventByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := controller.service.DeleteEventByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

// @PUT("/:id")
func (controller EventController) updateEventByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var m models.EventModel
	err2 := decoder.Decode(&m)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	response, err3 := controller.service.UpdateEventByID(id, m)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

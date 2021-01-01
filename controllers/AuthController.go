package controllers

import (
	"climbpass-matching-service/constants"
	"climbpass-matching-service/models"
	"climbpass-matching-service/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddAuthControllers adds controllers for auth
func AddAuthControllers(r *mux.Router) {
	controller := NewAuthController()
	baseRoute := constants.APIBasePath + "/auth"

	r.HandleFunc(baseRoute, controller.authenticateUser).Methods("POST")
	r.HandleFunc(baseRoute+"/create", controller.createNewAuthUser).Methods("POST")
}

// IAuthController interface for struct
type IAuthController interface {
	authenticateUser(w http.ResponseWriter, r *http.Request)
	createNewAuthUser(w http.ResponseWriter, r *http.Request)
}

// AuthController struct
type AuthController struct {
	service services.IAuthService
}

// NewAuthController does init for auth
func NewAuthController() IAuthController {
	service := services.NewAuthService()
	return AuthController{service}
}

// @POST("/")
func (controller AuthController) authenticateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var auth models.AuthUser
	err := decoder.Decode(&auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := controller.service.AuthenticateUser(auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}

// @POST("/create")
func (controller AuthController) createNewAuthUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var g models.AuthUser
	err := decoder.Decode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := controller.service.CreateAuth(g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}

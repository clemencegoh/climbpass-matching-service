package main

import (
	"climbpass-matching-service/constants"
	"climbpass-matching-service/controllers"
	"climbpass-matching-service/middlewares"
	"climbpass-matching-service/models"
	"climbpass-matching-service/repositories"
	"fmt"
	"gorm.io/gorm"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRouter specified all routes
func InitRouter(env string) *mux.Router {
	var db *gorm.DB

	if env == "test" {
		// Setup test DB
		db = repositories.ConnectSQLite()
	} else {
		// Setup prod DB
		db = repositories.Connect()
	}

	// Migrate schema
	err1 := db.AutoMigrate(
		&models.User{},
		&models.GymProfile{},
		&models.AuthUser{},
	)
	if err1 != nil {
		panic("failed to migrate models")
	}

	r := mux.NewRouter()

	// Serve static files in assets
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	// Routes
	r.HandleFunc(constants.APIBasePath+"/health", healthCheckHandler).Methods("GET")
	controllers.AddGymControllers(r)
	controllers.AddAuthControllers(r)

	// Middlewares
	r.Use(middlewares.ContentType)
	r.Use(middlewares.AuthMiddleware)

	return r
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

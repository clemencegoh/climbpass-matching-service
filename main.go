package main

import (
	"climbpass-matching-service/models"
	"climbpass-matching-service/repositories"
	"net/http"
)

func main() {

	// Setup DB
	db := repositories.Connect()

	// Migrate schema
	err1 := db.AutoMigrate(&models.GymModel{})
	if err1 != nil {
		panic("failed to migrate models")
	}

	// Listen and Serve
	err := http.ListenAndServe(":8080", InitRouter())
	if err != nil {
		panic(err)
	}
}

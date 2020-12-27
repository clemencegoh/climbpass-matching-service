package main

import (
	"climbpass-matching-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {

	setupDB()

	// Listen and Serve
	http.ListenAndServe(":8080", InitRouter())
}

func setupDB() {
	// setup local db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	// Migrate schema
	db.AutoMigrate(&models.GymModel{})
}

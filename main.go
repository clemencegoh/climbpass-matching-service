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

	// Init data
	// db.Create(&models.GymModel{
	// 	ID:       1,
	// 	Name:     "init_name",
	// 	Location: "some location",
	// })

	// db.Create(&models.GymModel{
	// 	ID:       2,
	// 	Name:     "init_name_2",
	// 	Location: "some location 2",
	// })
}

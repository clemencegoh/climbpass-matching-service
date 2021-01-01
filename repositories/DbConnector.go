package repositories

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

// sync to prevent race condition
var once sync.Once

// Global singleton
var instance *gorm.DB

// Connect to postgres db instance
func Connect() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=postgres password=password dbname=climbpass port=5432 sslmode=disable TimeZone=Asia/Singapore"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect to db")
		}
		instance = db
	})
	return instance
}

// ConnectSQLite connects to SQLite db
func ConnectSQLite() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect to db")
		}
		instance = db
	})
	return instance
}

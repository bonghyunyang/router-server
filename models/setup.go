package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("router.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB.")
	}

	return db
}

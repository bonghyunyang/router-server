package main

import (
	"router-server/models"
	"router-server/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Admin{}, &models.Sensor{}, &models.SensorData{}, &models.SensorFile{})

	r := routes.SetupRoutes(db)
	r.Run()
}

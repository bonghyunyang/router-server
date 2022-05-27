package routes

import (
	"router-server/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/sensor_data/:id", controllers.FindSensorData)
	r.GET("/sensor_data", controllers.FindAllSensorData)
	r.POST("/sensor_data", controllers.CreateSensorData)
	// r.PATCH("/tasks/:id", controllers.UpdateSensorData)
	r.DELETE("sensor_data/:id", controllers.DeleteSensorData)
	return r
}

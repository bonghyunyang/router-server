package controllers

import (
	"net/http"
	"router-server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 단일 sensorData 조회
func FindSensorData(c *gin.Context) { //
	var sensor_data models.SensorData
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&sensor_data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "해당 데이터가 없습니다."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sensor_data})
}

// 모든 sensorData 조회
func FindAllSensorData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var sensor_data []models.SensorData
	db.Find(&sensor_data)

	c.JSON(http.StatusOK, gin.H{"data": sensor_data})
}

// sensorData 등록
func CreateSensorData(c *gin.Context) {
	type SensorDetail struct {
		SensorID string  `json:"sensorId"`
		Sensing  int64   `json:"sensing"`
		Data     float32 `json:"data"`
	}

	var dataList struct {
		Data []SensorDetail `json:"data"`
	}

	// Json 타입으로 들어왔는지 체크
	if err := c.BindJSON(&dataList); err != nil {
		panic(err)
	}

	// 해당 데이터 UUID를 가진 센서가 없을 시 새로운 센서 생성, 있을 시 센서 데이터만 입력
	for _, s := range dataList.Data {
		var sensorModel models.Sensor
		db := c.MustGet("db").(*gorm.DB)
		if err := db.Where("UUID = ?", s.SensorID).First(&sensorModel).Error; err != nil {
			createNewSensor := models.Sensor{UUID: s.SensorID}
			db.Create(&createNewSensor)
			sensorData := models.SensorData{SensorID: s.SensorID, SensingDate: s.Sensing, Data: s.Data, Status: 0}
			db.Create(&sensorData)
		} else {
			sensor_data := models.SensorData{SensorID: s.SensorID, SensingDate: s.Sensing, Data: s.Data, Status: 0}
			db := c.MustGet("db").(*gorm.DB)
			db.Create(&sensor_data)
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// sensorData 삭제
func DeleteSensorData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var sensorData models.SensorData
	if err := db.Where("id = ?", c.Param("id")).First(&sensorData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "해당 데이터가 없습니다."})
		return
	}

	db.Delete(&sensorData)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

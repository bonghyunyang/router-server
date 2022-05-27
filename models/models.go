package models

import (
	"time"
)

// Admin Model
type Admin struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	UserId        string    `json:"user_id"`
	Password      string    `json:"password"`
	LastLoginDate time.Time `json:"last_login_date"`
}

// Sensor Model
type Sensor struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	UUID      string    `json:"uuid"`
}

// SensorData Model
type SensorData struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	SensorID    string    `json:"sensor_id"` // 센서 UUID
	Sensor      Sensor    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Data        float32   `json:"data" sql:"type:decimal(10,2);"`
	SensingDate int64     `json:"sensing_date"`
	Status      int64     `json:"status"`
}

// SensorFile Model
type SensorFile struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	SensorID  uint      `json:"sensor_id"`
	Sensor    Sensor    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FilePath  string    `json:"file_path"`
}

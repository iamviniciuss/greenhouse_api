package domain

import "time"

type TemperatureRepository interface {
	FindSensorById(id string) (*Sensor, error)
	FindLastValue(sensor_id string) (*HumidityRepositoryDTO, error)
	Create(temperature *HumidityRepositoryDTO) (*HumidityRepositoryDTO, error)
	CreateSensor(sensor *Sensor) (*Sensor, error)
}

type HumidityRepositoryDTO struct {
	ID        string    `json:"_id" bson:"_id"`
	SensorID  string    `json:"sensor_id" bson:"sensor_id"`
	CreatedAt time.Time `json:"_created_at"`
	Value     int64     `json:"value"`
}

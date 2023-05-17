package domain

import "time"

type TemperatureRepository interface {
	FindSensorById(id string) (*Sensor, error)
	FindLastValue(sensor_id string) (*HumidityRepositoryDTO, error)
	ListAll() ([]*HumidityRepositoryDTO, error)
	Create(temperature *HumidityRepositoryDTO) (*HumidityRepositoryDTO, error)
	CreateSensor(sensor *Sensor) (*Sensor, error)
}

type HumidityRepositoryDTO struct {
	ID         string    `json:"_id" bson:"_id"`
	SensorID   string    `json:"sensor_id" bson:"sensor_id"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	Value      int64     `json:"value"`
	Percentage int64     `json:"percentage"`
}

func (hr *HumidityRepositoryDTO) CalculatePercentage() {
	// 4500 --> 100%
	// value -> x
	hr.Percentage = (hr.Value * int64(100)) / int64(4500)
}

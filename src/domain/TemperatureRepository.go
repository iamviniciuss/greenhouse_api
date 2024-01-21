package domain

import (
	"time"

	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
)

type TemperatureRepository interface {
	FindSensorById(id string) (*Sensor, error)
	FindLastValue(sensor_id string) (*TemperatureRepositoryDTO, error)
	ListAll() ([]*TemperatureRepositoryDTO, error)
	Create(temperature *TemperatureRepositoryDTO) (*TemperatureRepositoryDTO, error)
	CreateSensor(sensor *Sensor) (*Sensor, error)
}

type TemperatureRepositoryDTO struct {
	ID                 string    `json:"_id" bson:"_id"`
	SensorID           string    `json:"sensor_id" bson:"sensor_id"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	Value              float64   `json:"value"`
	Humidity           float64   `json:"humidity_value" bson:"humidity_value"`
	Percentage         float64   `json:"percentage"`
	ExponentialAverage []float64 `json:"exponential_average"`
	MovelAverage       []float64 `json:"movel_average"`
}

func (hr *TemperatureRepositoryDTO) CalculatePercentage() {
	hr.Percentage = (float64(hr.Value) * float64(100)) / float64(4500)

}

func (hr *TemperatureRepositoryDTO) CalculateExponentialAverage(readings []float64) {
	hr.ExponentialAverage = shared.CalculateExponentialAverage(readings, 8)
}

func (hr *TemperatureRepositoryDTO) CalculateMovingAverage(readings []float64) {
	hr.MovelAverage = shared.CalculateMovingAverage(readings, 8)
}

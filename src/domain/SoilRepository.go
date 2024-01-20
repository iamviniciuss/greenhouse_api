package domain

import "time"

type SoilRepository interface {
	FindSensorById(id string) (*Sensor, error)
	FindLastValue(sensor_id string) (*HumidityRepositoryDTO, error)
	ListAll() ([]*HumidityRepositoryDTO, error)
	Create(temperature *HumidityRepositoryDTO) (*HumidityRepositoryDTO, error)
	CreateSensor(sensor *Sensor) (*Sensor, error)
}

type HumidityRepositoryDTO struct {
	ID                 string    `json:"_id" bson:"_id"`
	SensorID           string    `json:"sensor_id" bson:"sensor_id"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	Value              int64     `json:"value"`
	Percentage         float64   `json:"percentage"`
	ExponentialAverage []float64 `json:"exponential_average"`
	MovelAverage       []float64 `json:"movel_average"`
}

func (hr *HumidityRepositoryDTO) CalculatePercentage() {
	hr.Percentage = (float64(hr.Value) * float64(100)) / float64(4095)

}

func (hr *HumidityRepositoryDTO) CalculateExponentialAverage(readings []float64) {
	hr.ExponentialAverage = CalculateExponentialAverage(readings, 8)
}

func (hr *HumidityRepositoryDTO) CalculateMovelAverage(readings []float64) {
	hr.MovelAverage = CalculateMovingAverage(readings, 8)
}

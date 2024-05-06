package repository

import (
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
	mockTestify "github.com/stretchr/testify/mock"
)

type SoilRepositoryMock struct {
	mockTestify.Mock
}

func NewSoilRepositoryMock() *SoilRepositoryMock {
	return &SoilRepositoryMock{}
}

func (st *SoilRepositoryMock) FindSensorById(sensor_id string) (*shared.Sensor, error) {
	args := st.Called()
	output := args.Get(0).(*shared.Sensor)
	return output, args.Error(1)
}

func (st *SoilRepositoryMock) FindLastValue(temperature_id string) (*repository.HumidityRepositoryDTO, error) {
	args := st.Called()
	output := args.Get(0).(*repository.HumidityRepositoryDTO)
	return output, args.Error(1)
}

func (st *SoilRepositoryMock) Create(temperature *repository.HumidityRepositoryDTO) (*repository.HumidityRepositoryDTO, error) {
	args := st.Called(temperature)
	output := args.Get(0).(*repository.HumidityRepositoryDTO)
	return output, args.Error(1)
}

func (st *SoilRepositoryMock) ListAll() ([]*repository.HumidityRepositoryDTO, error) {
	args := st.Called()
	output := args.Get(0).([]*repository.HumidityRepositoryDTO)
	return output, args.Error(1)
}

func (st *SoilRepositoryMock) CreateSensor(sensor *shared.Sensor) (*shared.Sensor, error) {
	args := st.Called(sensor)
	output := args.Get(0).(*shared.Sensor)
	return output, args.Error(1)
}

func (st *SoilRepositoryMock) RecordMetric(metric repository.MetricRepositoryDTO) (repository.MetricRepositoryDTO, error) {
	st.Called(metric)
	return repository.MetricRepositoryDTO{}, nil
}

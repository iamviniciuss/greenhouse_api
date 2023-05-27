package repository

import (
	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	mockTestify "github.com/stretchr/testify/mock"
)

type SoilRepositoryMock struct {
	mockTestify.Mock
}

func NewSoilRepositoryMock() *SoilRepositoryMock {
	return &SoilRepositoryMock{}
}

func (st *SoilRepositoryMock) FindSensorById(sensor_id string) (*domain.Sensor, error) {
	args := st.Called()
	output := args.Get(0).(*domain.Sensor)
	return output, args.Error(1)
}

func (st *SoilRepositoryMock) FindLastValue(temperature_id string) (*domain.HumidityRepositoryDTO, error) {
	args := st.Called()
	output := args.Get(0).(*domain.HumidityRepositoryDTO)
	return output, args.Error(1)
}

func (st *SoilRepositoryMock) Create(temperature *domain.HumidityRepositoryDTO) (*domain.HumidityRepositoryDTO, error) {
	args := st.Called(temperature)
	output := args.Get(0).(*domain.HumidityRepositoryDTO)
	return output, args.Error(1)
}

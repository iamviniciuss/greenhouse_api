package application

import (
	"time"

	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	repository "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/repository"

	"testing"

	"github.com/stretchr/testify/suite"
)

type WaterPumpTestSuite struct {
	suite.Suite
}

func TestWaterPumpTestSuite(t *testing.T) {
	suite.Run(t, &WaterPumpTestSuite{})
}

func (suite *WaterPumpTestSuite) Test_Should_Return_Command_To_Turn_On_Water_Pump() {
	env := &domain.Envoironment{
		ID:           "1",
		Name:         "Humidity",
		GreenhouseID: "1",
	}

	sensor := &domain.Sensor{
		ID:            "1",
		Envoironments: env,
		Actuator: &domain.Actuator{
			ID:            "6",
			Name:          "Water Pump",
			Envoironments: env,
			Sensor:        &domain.Sensor{},
		},
		Name:         "Humidity Soil 01",
		GreenhouseID: "1",
		IdealValue:   []int{20, 26},
	}

	repository := repository.NewSoilRepositoryMock()
	repository.On("FindLastValue").Return(&domain.HumidityRepositoryDTO{
		ID:          "123",
		SensorID:    "1",
		CreatedAt:   time.Now().UTC(),
		Temperature: 16,
	}, nil)

	repository.On("FindSensorById").Return(sensor, nil)

	useCase := NewManageWaterPump(repository)

	output, err := useCase.Execute(&domain.Greenhouse{
		ID:      "1",
		Name:    "Saint's House",
		Sensors: []*domain.Sensor{sensor},
	})

	suite.Nil(err)
	suite.Equal(true, output.TurnOnWaterPump)
}

func (suite *WaterPumpTestSuite) Test_Should_Return_Command_To_Turn_On_Water_Pump_When_Humidity_Is_More_Than_Of_Ideal() {
	env := &domain.Envoironment{
		ID:           "1",
		Name:         "Humidity",
		GreenhouseID: "1",
	}

	sensor := &domain.Sensor{
		ID:            "1",
		Envoironments: env,
		Actuator: &domain.Actuator{
			ID:            "6",
			Name:          "Water Pump",
			Envoironments: env,
			Sensor:        &domain.Sensor{},
		},
		Name:         "Humidity Soil 01",
		GreenhouseID: "1",
		IdealValue:   []int{20, 26},
	}

	repository := repository.NewSoilRepositoryMock()
	repository.On("FindLastValue").Return(&domain.HumidityRepositoryDTO{
		ID:          "123",
		SensorID:    "1",
		CreatedAt:   time.Now().UTC(),
		Temperature: 27,
	}, nil)

	repository.On("FindSensorById").Return(sensor, nil)

	useCase := NewManageWaterPump(repository)

	output, err := useCase.Execute(&domain.Greenhouse{
		ID:      "1",
		Name:    "Saint's House",
		Sensors: []*domain.Sensor{sensor},
	})

	suite.Nil(err)
	suite.Equal(false, output.TurnOnWaterPump)
}

func (suite *WaterPumpTestSuite) AfterTest(suiteName, testName string) {}

func ManageTestTemperatureTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, &WaterPumpTestSuite{})
}

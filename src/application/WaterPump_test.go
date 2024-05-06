package application

import (
	"time"

	repo "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	repository_domain "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
	repository "github.com/iamviniciuss/greenhouse_api/src/infra/repository"

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WaterPumpTestSuite struct {
	suite.Suite
}

func TestWaterPumpTestSuite(t *testing.T) {
	suite.Run(t, &WaterPumpTestSuite{})
}

func (suite *WaterPumpTestSuite) Test_Should_Return_Command_To_Turn_On_Water_Pump() {
	env := &shared.Envoironment{
		ID:           "1",
		Name:         "Humidity",
		GreenhouseID: "1",
	}

	sensor := &shared.Sensor{
		ID:            "1",
		Envoironments: env,
		Actuator: &shared.Actuator{
			ID:            "6",
			Name:          "Water Pump",
			Envoironments: env,
			Sensor:        &shared.Sensor{},
		},
		Name:         "Humidity Soil 01",
		GreenhouseID: "1",
		IdealValue:   []int{20, 26},
	}

	repository := repository.NewSoilRepositoryMock()
	repository.On("FindLastValue").Return(&repository_domain.HumidityRepositoryDTO{
		ID:        "123",
		SensorID:  "1",
		CreatedAt: time.Now().UTC(),
		Value:     16,
	}, nil)

	repository.On("FindSensorById").Return(sensor, nil)
	repository.On("RecordMetric", mock.Anything).Maybe()

	useCase := NewManageWaterPump(repository)

	output, err := useCase.Execute(&shared.Greenhouse{
		ID:      "1",
		Name:    "Saint's House",
		Sensors: []*shared.Sensor{sensor},
	})

	suite.Nil(err)
	suite.Equal(true, output.TurnOnWaterPump)
}

func (suite *WaterPumpTestSuite) Test_Should_Return_Command_To_Turn_On_Water_Pump_When_Humidity_Is_More_Than_Of_Ideal() {
	env := &shared.Envoironment{
		ID:           "1",
		Name:         "Humidity",
		GreenhouseID: "1",
	}

	sensor := &shared.Sensor{
		ID:            "1",
		Envoironments: env,
		Actuator: &shared.Actuator{
			ID:            "6",
			Name:          "Water Pump",
			Envoironments: env,
			Sensor:        &shared.Sensor{},
		},
		Name:         "Humidity Soil 01",
		GreenhouseID: "1",
		IdealValue:   []int{20, 26},
	}

	repository := repository.NewSoilRepositoryMock()
	repository.On("FindLastValue").Return(&repository_domain.HumidityRepositoryDTO{
		ID:        "123",
		SensorID:  "1",
		CreatedAt: time.Now().UTC(),
		Value:     27,
	}, nil)

	repository.On("FindSensorById").Return(sensor, nil)
	repository.On("RecordMetric", mock.Anything).Return(repo.MetricRepositoryDTO{}, nil)

	useCase := NewManageWaterPump(repository)

	output, err := useCase.Execute(&shared.Greenhouse{
		ID:      "1",
		Name:    "Saint's House",
		Sensors: []*shared.Sensor{sensor},
	})

	suite.Nil(err)
	suite.Equal(false, output.TurnOnWaterPump)
}

func (suite *WaterPumpTestSuite) AfterTest(suiteName, testName string) {}

func ManageTestTemperatureTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, &WaterPumpTestSuite{})
}

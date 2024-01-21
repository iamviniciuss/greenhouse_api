package domain

import (
	"time"

	"testing"

	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
	"github.com/stretchr/testify/suite"
)

type ManageTemperatureTestSuite struct {
	suite.Suite
}

func TestManageTemperatureTestSuite(t *testing.T) {
	suite.Run(t, &ManageTemperatureTestSuite{})
}

func (suite *ManageTemperatureTestSuite) Test_Should_Return_Command_To_Turn_Off_Water_Pump() {
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

	humidity := &repository.HumidityRepositoryDTO{
		ID:        "123",
		SensorID:  "1",
		CreatedAt: time.Now().UTC(),
		Value:     24,
	}

	waterPump := NewWaterPumpState()
	command, err := waterPump.ManageState(sensor, humidity)

	suite.Nil(err)
	suite.Equal(false, command.TurnOnWaterPump)
}

func (suite *ManageTemperatureTestSuite) Test_Should_Return_Command_To_Turn_On_Water_Pump() {
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

	humidity := &repository.HumidityRepositoryDTO{
		ID:        "123",
		SensorID:  "1",
		CreatedAt: time.Now().UTC(),
		Value:     19,
	}

	waterPump := NewWaterPumpState()
	command, err := waterPump.ManageState(sensor, humidity)

	suite.Nil(err)
	suite.Equal(true, command.TurnOnWaterPump)
}

func (suite *ManageTemperatureTestSuite) Test_Should_Return_Command_To_Turn_On_Water_Pump_When_Humidity_Is_More_Than_Of_Ideal() {
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

	humidity := &repository.HumidityRepositoryDTO{
		ID:        "123",
		SensorID:  "1",
		CreatedAt: time.Now().UTC(),
		Value:     27,
	}

	waterPump := NewWaterPumpState()
	command, err := waterPump.ManageState(sensor, humidity)

	suite.Nil(err)
	suite.Equal(false, command.TurnOnWaterPump)
}

func (suite *ManageTemperatureTestSuite) AfterTest(suiteName, testName string) {}

func ManageTestTemperatureTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, &ManageTemperatureTestSuite{})
}

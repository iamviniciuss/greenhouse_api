package domain

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnergyConsumptionTestSuite struct {
	suite.Suite
}

func TestEnergyConsumptionTestSuite(t *testing.T) {
	suite.Run(t, &EnergyConsumptionTestSuite{})
}

func (suite *EnergyConsumptionTestSuite) Test_Calculate_The_Producing_Of_Water_Sended_To_Greenhouse() {
	waterPump := WaterPump{
		ImpulseWaterPerMinutes: 1.5,
		Voltage:                5,
		ElectricCurrent:        0.2,
	}

	output := waterPump.EnergyConsumption(1)

	suite.Equal(1.0, output.EnergyConsumption, "consumption in hours of the water pump")
}

func (suite *EnergyConsumptionTestSuite) Test_Calculate_The_Producing_Of_Water_Sended_To_Greenhouse_2() {
	waterPump := WaterPump{
		ImpulseWaterPerMinutes: 1.5,
		Voltage:                5,
		ElectricCurrent:        0.2,
	}
	output := waterPump.EnergyConsumption(3.5)

	suite.Equal(3.5, output.EnergyConsumption, "consumption in hours of the water pump")
}

func (suite *EnergyConsumptionTestSuite) Test_Calculate_The_CalculatePower() {
	waterPump := WaterPump{
		ImpulseWaterPerMinutes: 1.5,
		Voltage:                5,
		ElectricCurrent:        0.2,
	}
	output := waterPump.CalculatePower()

	suite.Equal(1.0, output, "water pump power")
}

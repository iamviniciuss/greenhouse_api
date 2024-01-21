package domain

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProducingWaterTestSuite struct {
	suite.Suite
}

func TestProducingWaterTestSuite(t *testing.T) {
	suite.Run(t, &ProducingWaterTestSuite{})
}

func (suite *ProducingWaterTestSuite) Test_Calculate_The_Producing_Of_Water_Sended_To_Greenhouse() {
	waterPump := WaterPump{
		ImpulseWaterPerMinutes: 1.5,
	}
	output := waterPump.ProducingWater(1)

	suite.Equal(1.5, output.NumberOfLitersPumped, "1 minute on should correspond to 1.5 liters sent to the garden")
}

func (suite *ProducingWaterTestSuite) Test_Calculate_The_Producing_Of_Water_Sended_To_Greenhouse_2() {
	waterPump := WaterPump{
		ImpulseWaterPerMinutes: 1.5,
	}
	output := waterPump.ProducingWater(1.5)

	suite.Equal(2.25, output.NumberOfLitersPumped, "1 minute on must correspond to 2.25 liters sent to the garden")
}

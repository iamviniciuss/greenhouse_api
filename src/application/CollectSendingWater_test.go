package application

import (
	"time"

	repo "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
	domain "github.com/iamviniciuss/greenhouse_api/src/domain/soil"
	repository "github.com/iamviniciuss/greenhouse_api/src/infra/repository"

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CollectSendingWaterTestSuite struct {
	suite.Suite
}

func TestCollectSendingWaterTestSuite(t *testing.T) {
	suite.Run(t, &CollectSendingWaterTestSuite{})
}

func (suite *CollectSendingWaterTestSuite) Test_Sould_Record_The_Metrics_Of_Water_Pump() {
	repository := repository.NewSoilRepositoryMock()

	repository.On("RecordMetric", mock.Anything).Return(repo.MetricRepositoryDTO{}, nil)
	startDate := time.Now()

	output, err := NewCollectSendingWater(repository).Execute(domain.WaterPump{
		ImpulseWaterPerMinutes: 1.5,
		Voltage:                5,
		ElectricCurrent:        0.2,
	}, shared.Event{
		ID:      "1",
		Started: startDate,
		Ended:   startDate.Add(3 * time.Minute),
	})

	repository.AssertNumberOfCalls(suite.T(), "RecordMetric", 2)
	suite.Nil(err)
	suite.NotNil(output)
	suite.Equal(output.EnergyConsume, 3.0)
	suite.Equal(output.WaterBombed, 4.5)
}

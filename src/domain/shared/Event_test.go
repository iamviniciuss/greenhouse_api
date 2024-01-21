package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type EventTestSuite struct {
	suite.Suite
}

func TestEventTestSuite(t *testing.T) {
	suite.Run(t, &EventTestSuite{})
}

func (suite *EventTestSuite) Test_Should_Calculate_The_Difference_between_Start_and_End_Date_of_event() {
	startDate := time.Now()

	event := Event{
		Started: startDate,
		Ended:   startDate.Add(3 * time.Minute),
	}

	suite.Equal(3.0, event.Duration())
}

func (suite *EventTestSuite) Test_Should_Calculate_The_Difference_between_Start_and_End_Date_of_event_2() {
	startDate := time.Now()

	event := Event{
		Started: startDate,
		Ended:   startDate,
	}

	suite.Equal(0.0, event.Duration())
}

func (suite *EventTestSuite) Test_Should_Calculate_The_Difference_between_Start_and_End_Date_of_event_3() {
	startDate := time.Now()

	event := Event{
		Started: startDate,
		Ended:   startDate.Add(1 * time.Minute),
	}

	suite.Equal(1.0, event.Duration())
}

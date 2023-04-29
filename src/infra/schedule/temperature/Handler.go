package infra

import (
	"context"
	"log"
	"time"

	"github.com/procyon-projects/chrono"
)

type HandleTemperatureInput struct {
}

type HandleTemperature struct {
}

func NewHandleTemperature() *HandleTemperature {
	return &HandleTemperature{}
}

func (se *HandleTemperature) Handler() error {

	taskScheduler := chrono.NewDefaultTaskScheduler()

	_, err := taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		log.Print("Fixed Rate of 5 seconds")
	}, 5*time.Second)

	if err == nil {
		log.Print("Task has been scheduled successfully.")
	}

	return nil
}

package domain

import (
	"fmt"

	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
)

type WaterPump struct {
	ImpulseWaterPerMinutes float64
	Voltage                float64
	ElectricCurrent        float64
}

type WaterPumpState struct {
	TurnOnWaterPump bool
}

func NewWaterPumpState() *WaterPump {
	return &WaterPump{}
}

func (wp *WaterPump) ManageState(sensor *shared.Sensor, humidity *repository.HumidityRepositoryDTO) (*WaterPumpState, error) {
	itHumidityOk := shared.IsInRange(
		int(humidity.Value),
		int(sensor.IdealValue[0]),
		int(sensor.IdealValue[1]),
	)

	fmt.Println("\n\nWaterPump- ManageState:", sensor.IdealValue[0], sensor.IdealValue[1])

	if itHumidityOk {
		fmt.Println("Drop a bomb! Humidity is at the intermediate level.")
		return &WaterPumpState{
			TurnOnWaterPump: false,
		}, nil
	}

	itIsHighHumidity := shared.HumidityIsHigh(int(humidity.Value), int(sensor.IdealValue[1]))
	if itIsHighHumidity {
		fmt.Println("High Humidity! Turn off water pump!")
		return &WaterPumpState{
			TurnOnWaterPump: false,
		}, nil
	}

	fmt.Println("LOW humidity!! turn on the pump")

	return &WaterPumpState{
		TurnOnWaterPump: true,
	}, nil
}

package domain

import "fmt"

type Temperature struct {
}

type TemperatureState struct {
	Command string `json:"command"`
}

func NewTemperatureState() *Temperature {
	return &Temperature{}
}

func (wp *Temperature) ManageState(sensor *Sensor, humidity *TemperatureRepositoryDTO) (*TemperatureState, error) {
	itTemperatureOk := isInRange(
		int(humidity.Value)+(int(sensor.IdealValue[0])/int(sensor.IdealValue[1])),
		int(sensor.IdealValue[0]),
		int(sensor.IdealValue[1]),
	)

	if itTemperatureOk {
		fmt.Println("Turn off All atuators! The temperature is at the intermediate level.")
		return &TemperatureState{
			Command: "TURN_OFF_ALL",
		}, nil
	}

	itIsHighTemperature := humidityIsHigh(int(humidity.Value), int(sensor.IdealValue[1]))
	if itIsHighTemperature {
		fmt.Println("High temperature! Turn on 	COOLER!")
		return &TemperatureState{
			Command: "TURN_ON_COOLER",
		}, nil
	}

	fmt.Println("LOW temperature!! turn on LIGHT")

	return &TemperatureState{
		Command: "TURN_ON_LIGHT",
	}, nil
}

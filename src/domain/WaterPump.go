package domain

type WaterPump struct {
}

type WaterPumpState struct {
	TurnOnWaterPump bool
}

func NewWaterPumpState() *WaterPump {
	return &WaterPump{}
}

func (wp *WaterPump) ManageState(sensor *Sensor, humidity *HumidityRepositoryDTO) (*WaterPumpState, error) {
	itHumidityOk := isInRange(
		int(humidity.Value)+(int(sensor.IdealValue[0])/int(sensor.IdealValue[1])),
		int(sensor.IdealValue[0]),
		int(sensor.IdealValue[1]),
	)

	if itHumidityOk {
		// fmt.Println("Drop a bomb! Humidity is at the intermediate level.")
		return &WaterPumpState{
			TurnOnWaterPump: false,
		}, nil
	}

	itIsHighHumidity := humidityIsHigh(int(humidity.Value), int(sensor.IdealValue[1]))
	if itIsHighHumidity {
		// fmt.Println("High Humidity! Turn off water pump!")
		return &WaterPumpState{
			TurnOnWaterPump: false,
		}, nil
	}

	// fmt.Println("LOW humidity!! turn on the pump")

	return &WaterPumpState{
		TurnOnWaterPump: true,
	}, nil
}

func isInRange(num, min, max int) bool {
	return num >= min && num <= max
}

func humidityIsHigh(num, max int) bool {
	return num > max
}

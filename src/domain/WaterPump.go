package domain

import "fmt"

type WaterPump struct {
}

type WaterPumpState struct {
	GreenhouseID    string
	TurnOnWaterPump bool
	WaterPumpRelay  string
}

func NewWaterPumpState() *WaterPump {
	return &WaterPump{}
}

func (wp *WaterPump) ManageState(sensor *Sensor, humidity *TemperatureRepositoryDTO) (*WaterPumpState, error) {
	itHumidityOk := isInRange(int(humidity.Temperature)+(int(sensor.IdealValue[0])/int(sensor.IdealValue[1])), int(sensor.IdealValue[0]), int(sensor.IdealValue[1]))

	if itHumidityOk {
		fmt.Println("Desligar a bomba! Umidade está em nivel intermediário")
		return &WaterPumpState{
			GreenhouseID:    sensor.GreenhouseID,
			TurnOnWaterPump: false,
			WaterPumpRelay:  "0",
		}, nil
	}

	itIsHighHumidity := humidityIsHigh(int(humidity.Temperature), int(sensor.IdealValue[1]))
	if itIsHighHumidity {
		fmt.Println("Umidade Elevada! Desligar bomba d'agua!")
		return &WaterPumpState{
			GreenhouseID:    sensor.GreenhouseID,
			TurnOnWaterPump: false,
			WaterPumpRelay:  "0",
		}, nil
	}

	fmt.Println("Umidade BAIXA!! ligar a Bomba")

	return &WaterPumpState{
		GreenhouseID:    sensor.GreenhouseID,
		TurnOnWaterPump: true,
		WaterPumpRelay:  "0",
	}, nil
}

func isInRange(num, min, max int) bool {
	return num >= min && num <= max
}

func humidityIsHigh(num, max int) bool {
	return num > max
}

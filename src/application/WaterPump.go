package application

import (
	"fmt"

	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
)

type ManageWaterPump struct {
	domain.TemperatureRepository
}

type ManageWaterPumpOutput struct {
	GreenhouseID    string
	TurnOnWaterPump bool
	WaterPumpRelay  string
}

func NewManageWaterPump(tp domain.TemperatureRepository) *ManageWaterPump {
	return &ManageWaterPump{
		TemperatureRepository: tp,
	}
}

func (m *ManageWaterPump) Execute(greenhouse *domain.Greenhouse) (*ManageWaterPumpOutput, error) {

	if len(greenhouse.Sensors) == 0 {
		return nil, fmt.Errorf("there aren't any sensors")
	}

	sensorID := greenhouse.Sensors[0].ID

	fmt.Println("sensorID:", sensorID)

	humidity, err := m.TemperatureRepository.FindLastValue(sensorID)
	if err != nil {
		fmt.Println("not found humidity:")
		return nil, err
	}

	sensor, err := m.TemperatureRepository.FindSensorById(humidity.SensorID)
	if err != nil {
		fmt.Println("not found sensor:")
		return nil, err
	}

	output, err := domain.NewWaterPumpState().ManageState(sensor, humidity)
	if err != nil {
		fmt.Println("not found NewWaterPumpState:")
		return nil, err
	}

	return &ManageWaterPumpOutput{
		GreenhouseID:    sensor.GreenhouseID,
		TurnOnWaterPump: output.TurnOnWaterPump,
		WaterPumpRelay:  "0",
	}, nil

}

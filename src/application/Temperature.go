package application

import (
	"encoding/json"
	"fmt"
	"os"

	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
	temperature_domain "github.com/iamviniciuss/greenhouse_api/src/domain/temperature"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type ManageTemperature struct {
	repository.TemperatureRepository
}

type ManageTemperatureOutput struct {
	GreenhouseID   string
	Command        string
	WaterPumpRelay string
}

func NewManageTemperature(tp repository.TemperatureRepository) *ManageTemperature {
	return &ManageTemperature{
		TemperatureRepository: tp,
	}
}

func (m *ManageTemperature) Execute(greenhouse *shared.Greenhouse) (*ManageTemperatureOutput, error) {

	if len(greenhouse.Sensors) == 0 {
		return nil, fmt.Errorf("there aren't any sensors")
	}

	sensorID := greenhouse.Sensors[0].ID

	fmt.Println("sensorID:", sensorID)

	temperature, err := m.TemperatureRepository.FindLastValue(sensorID)
	if err != nil {
		fmt.Println("not found temperature:")
		return nil, err
	}

	sensor, err := m.TemperatureRepository.FindSensorById(temperature.SensorID)
	if err != nil {
		fmt.Println("not found sensor:", temperature.SensorID)
		return nil, err
	}

	output, err := temperature_domain.NewTemperatureState().ManageState(sensor, temperature)
	if err != nil {
		fmt.Println("not found NewTemperatureState:")
		return nil, err
	}

	return &ManageTemperatureOutput{
		GreenhouseID:   sensor.GreenhouseID,
		Command:        output.Command,
		WaterPumpRelay: "0",
	}, nil

}

func (m *ManageTemperature) GetCommand(mqttClient MQTT.Client) (*TemperatureCtrlOutput, error) {
	fmt.Println("ManageTemperature - Executing GetCommand")

	greenhouse := &shared.Greenhouse{
		ID:   "01",
		Name: "ESP32_HOUSE_VINICIUS",
		Sensors: []*shared.Sensor{
			&shared.Sensor{
				ID:            "645d82f4d2d163d2edc380a5",
				Envoironments: &shared.Envoironment{},
				Actuator: &shared.Actuator{
					ID:   "1",
					Name: "Bomba d'água",
				},
				IdealValue: []int{1400, 1450},
				Name:       "FC-28 - Sensor de umidade do solo",
			},
		},
	}

	res, err := m.Execute(greenhouse)

	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(TemperatureCtrlOutput{
		Command: res.Command,
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err

	}

	token := mqttClient.Publish(os.Getenv("TEMPERATURE_PUBLISHER")+greenhouse.Sensors[0].ID, 0, false, payload)

	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error().Error())
		return nil, token.Error()
	}

	return &TemperatureCtrlOutput{
		Command: res.Command,
	}, nil

}

type TemperatureCtrlOutput struct {
	Command string `json:"command"`
}

package application

import (
	"encoding/json"
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	domain "github.com/iamviniciuss/greenhouse_api/src/domain"
)

type ManageWaterPump struct {
	domain.SoilRepository
}

type ManageWaterPumpOutput struct {
	GreenhouseID    string
	TurnOnWaterPump bool
	WaterPumpRelay  string
}

func NewManageWaterPump(tp domain.SoilRepository) *ManageWaterPump {
	return &ManageWaterPump{
		SoilRepository: tp,
	}
}

func (m *ManageWaterPump) Execute(greenhouse *domain.Greenhouse) (*ManageWaterPumpOutput, error) {

	if len(greenhouse.Sensors) == 0 {
		return nil, fmt.Errorf("there aren't any sensors")
	}

	sensorID := greenhouse.Sensors[0].ID

	// fmt.Println("sensorID:", sensorID)

	humidity, err := m.SoilRepository.FindLastValue(sensorID)
	if err != nil {
		// fmt.Println("not found humidity:")
		return nil, err
	}

	sensor, err := m.SoilRepository.FindSensorById(humidity.SensorID)
	if err != nil {
		// fmt.Println("not found sensor:", humidity.SensorID)
		return nil, err
	}

	output, err := domain.NewWaterPumpState().ManageState(sensor, humidity)
	if err != nil {
		// fmt.Println("not found NewWaterPumpState:")
		return nil, err
	}

	return &ManageWaterPumpOutput{
		GreenhouseID:    sensor.GreenhouseID,
		TurnOnWaterPump: output.TurnOnWaterPump,
		WaterPumpRelay:  "0",
	}, nil

}

func (m *ManageWaterPump) GetCommand(mqttClient MQTT.Client) (*WaterPumpCtrlOutput, error) {
	fmt.Println("Executing GetCommand....")

	greenhouse := &domain.Greenhouse{
		ID:   "01",
		Name: "ESP32_HOUSE_VINICIUS",
		Sensors: []*domain.Sensor{
			{
				ID:            "645d82f4d2d163d2edc380a5",
				Envoironments: &domain.Envoironment{},
				Actuator: &domain.Actuator{
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

	payload, err := json.Marshal(WaterPumpCtrlOutput{
		TurnOnWaterPump: res.TurnOnWaterPump,
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err

	}

	token := mqttClient.Publish(os.Getenv("WATER_PUMP_PUBLISHER")+greenhouse.Sensors[0].ID, 0, false, payload)

	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error().Error())
		return nil, token.Error()
	}

	return &WaterPumpCtrlOutput{
		TurnOnWaterPump: res.TurnOnWaterPump,
	}, nil

}

type WaterPumpCtrlOutput struct {
	TurnOnWaterPump bool `json:"turnOn"`
}

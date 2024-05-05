package application

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
	domain "github.com/iamviniciuss/greenhouse_api/src/domain/soil"
	soil_domain "github.com/iamviniciuss/greenhouse_api/src/domain/soil"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type ManageWaterPump struct {
	repository.SoilRepository
}

type ManageWaterPumpOutput struct {
	GreenhouseID    string
	TurnOnWaterPump bool
	WaterPumpRelay  string
}

func NewManageWaterPump(tp repository.SoilRepository) *ManageWaterPump {
	return &ManageWaterPump{
		SoilRepository: tp,
	}
}

func (m *ManageWaterPump) Execute(greenhouse *shared.Greenhouse) (*ManageWaterPumpOutput, error) {

	if len(greenhouse.Sensors) == 0 {
		return nil, fmt.Errorf("there aren't any sensors")
	}

	sensorID := greenhouse.Sensors[0].ID

	humidity, err := m.SoilRepository.FindLastValue(sensorID)
	if err != nil {
		return nil, fmt.Errorf("not found last value to humidity")
	}

	sensor, err := m.SoilRepository.FindSensorById(humidity.SensorID)
	if err != nil {
		return nil, fmt.Errorf("not found sensor by id")
	}

	output, err := soil_domain.NewWaterPumpState().ManageState(sensor, humidity)
	if err != nil {
		return nil, fmt.Errorf("error to make command to water pump")
	}

	Started := time.Now().Add(-(10 * time.Minute))
	Ended := time.Now().Add(2 * time.Minute)

	NewCollectSendingWater(m.SoilRepository).Execute(domain.WaterPump{
		ImpulseWaterPerMinutes: 1.5,
		Voltage:                5,
		ElectricCurrent:        0.2,
	}, shared.Event{
		ID:      "",
		Started: Started,
		Ended:   Ended,
	})
	return &ManageWaterPumpOutput{
		GreenhouseID:    sensor.GreenhouseID,
		TurnOnWaterPump: output.TurnOnWaterPump,
		WaterPumpRelay:  "0",
	}, nil

}

func (m *ManageWaterPump) GetCommand(mqttClient MQTT.Client) (*WaterPumpCtrlOutput, error) {
	fmt.Println("Executing GetCommand....")

	greenhouse := &shared.Greenhouse{
		ID:   "01",
		Name: "ESP32_HOUSE_VINICIUS",
		Sensors: []*shared.Sensor{
			{
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

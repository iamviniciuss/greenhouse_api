package broker

import (
	"encoding/json"
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	application "github.com/iamviniciuss/greenhouse_api/src/application"
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
)

type WaterPumpTopicoCommand struct {
	humidityRepository repository.SoilRepository
	mqttClient         MQTT.Client
	Topico             string
}

func NewWaterPumpTopicoCommand(humidityRepository repository.SoilRepository, client MQTT.Client, topic string) *WaterPumpTopicoCommand {
	return &WaterPumpTopicoCommand{
		Topico:             topic,
		humidityRepository: humidityRepository,
		mqttClient:         client,
	}
}

func (c *WaterPumpTopicoCommand) SetMQTTClient(mqttClient MQTT.Client) {
	c.mqttClient = mqttClient
}

func (c *WaterPumpTopicoCommand) Execute(currentTopic string, message []byte) {
	if currentTopic != os.Getenv("WATER_PUMP_SUBSCRIBE") {
		return
	}

	var inputJSON RegisterHumidityCtrlnput
	err := json.Unmarshal(message, &inputJSON)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Humidity: ", inputJSON.Humidity)
	_, err = c.humidityRepository.Create(&repository.HumidityRepositoryDTO{
		SensorID: inputJSON.SensorID,
		Value:    inputJSON.Humidity,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go application.NewManageWaterPump(c.humidityRepository).GetCommand(c.mqttClient)

	fmt.Println("Executando ação para o WaterPumpTopicoCommand")
}

func (c *WaterPumpTopicoCommand) GetTopic() string {
	return c.Topico
}

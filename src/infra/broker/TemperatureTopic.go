package broker

import (
	"encoding/json"
	"fmt"
	"os"

	application "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/application"
	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type TemperatureTopicoCommand struct {
	temperatureRepository domain.TemperatureRepository
	mqttClient            MQTT.Client
	Topico                string
}

func NewTemperatureTopicoCommand(temperatureRepository domain.TemperatureRepository, client MQTT.Client, topic string) *TemperatureTopicoCommand {
	return &TemperatureTopicoCommand{
		Topico:                topic,
		temperatureRepository: temperatureRepository,
		mqttClient:            client,
	}
}

func (c *TemperatureTopicoCommand) SetMQTTClient(mqttClient MQTT.Client) {
	c.mqttClient = mqttClient
}

func (c *TemperatureTopicoCommand) Execute(currentTopic string, message []byte) {

	if currentTopic != os.Getenv("TEMPERATURE_SUBSCRIBE") {
		return
	}

	fmt.Println("Executando ação para o TemperatureTopicoCommand")

	var inputJSON RegisterTemperatureCtrlnput
	err := json.Unmarshal(message, &inputJSON)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Temperature: ", inputJSON.Temperature)
	_, err = c.temperatureRepository.Create(&domain.TemperatureRepositoryDTO{
		SensorID: inputJSON.SensorID,
		Value:    inputJSON.Temperature,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go application.NewManageTemperature(c.temperatureRepository).GetCommand(c.mqttClient)

}

func (c *TemperatureTopicoCommand) GetTopic() string {
	return c.Topico
}

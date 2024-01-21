package broker

import (
	"encoding/json"
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	application "github.com/iamviniciuss/greenhouse_api/src/application"
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
)

type RegisterTemperature struct {
	Humidity    float64 `json:"humidity,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	SensorID    string  `json:"sensor_id,omitempty"`
}

type TemperatureTopicoCommand struct {
	temperatureRepository repository.TemperatureRepository
	mqttClient            MQTT.Client
	Topico                string
}

func NewTemperatureTopicoCommand(temperatureRepository repository.TemperatureRepository, client MQTT.Client, topic string) *TemperatureTopicoCommand {
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

	var inputJSON RegisterTemperature
	err := json.Unmarshal(message, &inputJSON)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Temperature: ", inputJSON.Temperature)
	_, err = c.temperatureRepository.Create(&repository.TemperatureRepositoryDTO{
		SensorID: inputJSON.SensorID,
		Value:    inputJSON.Temperature,
		Humidity: inputJSON.Humidity,
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

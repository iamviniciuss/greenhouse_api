package main

import (
	"os"

	mongodb "github.com/iamviniciuss/greenhouse_api/src/infra/database/mongodb"
	httpService "github.com/iamviniciuss/greenhouse_api/src/infra/http"
	"github.com/iamviniciuss/greenhouse_api/src/infra/repository"
	healthCheck "github.com/iamviniciuss/greenhouse_api/src/infra/web/healthcheck"
	humidity "github.com/iamviniciuss/greenhouse_api/src/infra/web/v1/humidity"
	sensor "github.com/iamviniciuss/greenhouse_api/src/infra/web/v1/sensor"
	temperature "github.com/iamviniciuss/greenhouse_api/src/infra/web/v1/temperature"
)

func main() {

	// hostname, _ := os.Hostname()
	http := httpService.NewFiberHttp()
	mongo := mongodb.NewMongoConnection()
	mongo.Info()

	soildRepository := repository.NewSoilRepositoryMongo(mongo)
	temperatureRepository := repository.NewTemperatureRepositoryMongo(mongo)

	healthCheck.HealthCheckRouter(http)
	humidity.HumidityRouter(http, soildRepository)
	temperature.TemperatureRouter(http, temperatureRepository)
	sensor.SensorRouter(http, soildRepository)

	// brokerClient := "esp32/greenhouse-" + hostname
	// fmt.Println("Broker Client:", brokerClient)
	// mqttBroker := broker.NewMQTTBroker(brokerClient)
	// topicsToConsume := broker.
	// 	NewTopicsToConsumer().
	// 	Add(broker.NewTemperatureTopicoCommand(temperatureRepository, mqttBroker.GetClient(), os.Getenv("TEMPERATURE_SUBSCRIBE"))).
	// 	Add(broker.NewWaterPumpTopicoCommand(soildRepository, mqttBroker.GetClient(), os.Getenv("WATER_PUMP_SUBSCRIBE")))

	// mqttBroker.SetSubscribeTopics(topicsToConsume)

	// go mqttBroker.StartConsumers()

	err := http.ListenAndServe(os.Getenv("PORT"))
	if err != nil {
		// mqttBroker.Disconnect()
		panic(err)
	}
	// mqttBroker.Disconnect()
	panic("**** Close app! *****")
}

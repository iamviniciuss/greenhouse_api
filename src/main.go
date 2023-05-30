package main

import (
	"os"

	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/broker"
	mongodb "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/database/mongodb"
	httpService "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/repository"
	healthCheck "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/healthcheck"
	humidity "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/v1/humidity"
	sensor "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/v1/sensor"
	temperature "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/v1/temperature"
)

func main() {

	http := httpService.NewFiberHttp()
	mongo := mongodb.NewMongoConnection()
	mongo.Info()

	soildRepository := repository.NewSoilRepositoryMongo(mongo)
	temperatureRepository := repository.NewTemperatureRepositoryMongo(mongo)

	healthCheck.HealthCheckRouter(http)
	humidity.HumidityRouter(http, soildRepository)
	temperature.TemperatureRouter(http, temperatureRepository)
	sensor.SensorRouter(http, soildRepository)

	mqttBroker := broker.NewMQTTBroker("sdk-nodejs-v2")
	topicsToConsume := broker.
		NewTopicsToConsumer().
		Add(broker.NewTemperatureTopicoCommand(temperatureRepository, mqttBroker.GetClient(), os.Getenv("TEMPERATURE_SUBSCRIBE"))).
		Add(broker.NewWaterPumpTopicoCommand(soildRepository, mqttBroker.GetClient(), os.Getenv("WATER_PUMP_SUBSCRIBE")))

	mqttBroker.SetSubscribeTopics(topicsToConsume)

	go mqttBroker.StartConsumers()

	err := http.ListenAndServe(os.Getenv("PORT"))
	if err != nil {
		mqttBroker.Disconnect()
		panic(err)
	}
	mqttBroker.Disconnect()
	panic("**** Close app! *****")
}

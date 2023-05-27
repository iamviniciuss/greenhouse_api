package main

import (
	"os"

	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/broker"
	mongodb "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/database/mongodb"
	httpService "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/repository"
	healthCheck "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/healthcheck"
	humidity "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/v1/humidity"
	temperature "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/v1/temperature"
)

func main() {

	http := httpService.NewFiberHttp()
	mongo := mongodb.NewMongoConnection()
	mongo.Info()

	soildRepo := repository.NewSoilRepositoryMongo(mongo)
	repo := repository.NewTemperatureRepositoryMongo(mongo)

	healthCheck.HealthCheckRouter(http)
	humidity.HumidityRouter(http, soildRepo)
	temperature.TemperatureRouter(http, repo)

	// temperature.NewHandleTemperature().Handler()

	mqttBroker := broker.NewMQTTBroker(soildRepo, os.Getenv("WATER_PUMP_SUBSCRIBE"))
	mqttClient := mqttBroker.MQTTClient()

	mqttBroker2 := broker.NewMQTTBroker(soildRepo, os.Getenv("TEMPERATURE_SUBSCRIBE"))
	mqttClient2 := mqttBroker2.MQTTClient()

	go mqttBroker.MQTTConsumer()
	go mqttBroker2.MQTTConsumer()

	err := http.ListenAndServe(os.Getenv("PORT"))
	if err != nil {
		mqttClient.Disconnect(250)
		mqttClient2.Disconnect(250)
		panic(err)
	}
	mqttClient.Disconnect(250)
	mqttClient2.Disconnect(250)
	panic("**** Close app! *****")
}

package main

import (
	"os"

	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/broker"
	mongodb "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/database/mongodb"
	httpService "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/repository"
	healthCheck "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/healthcheck"
	humidity "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/v1/humidity"
)

func main() {

	http := httpService.NewFiberHttp()
	mongo := mongodb.NewMongoConnection()
	mongo.Info()

	repo := repository.NewTemperatureRepositoryMongo(mongo)

	healthCheck.HealthCheckRouter(http)
	humidity.HumidityRouter(http, repo)

	// temperature.NewHandleTemperature().Handler()
	mqttBroker := broker.NewMQTTBroker(repo)
	mqttClient := mqttBroker.MQTTClient()

	go mqttBroker.MQTTConsumer()

	err := http.ListenAndServe(os.Getenv("PORT"))
	if err != nil {
		mqttClient.Disconnect(250)
		panic(err)
	}
	mqttClient.Disconnect(250)
	panic("**** Close app! *****")
}

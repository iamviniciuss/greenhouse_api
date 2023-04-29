package main

import (
	"os"

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

	err := http.ListenAndServe(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}

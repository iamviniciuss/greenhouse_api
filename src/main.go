package main

import (
	"os"

	mongodb "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/database/mongodb"
	httpService "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
	healthCheck "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/healthcheck"
)

func main() {

	http := httpService.NewFiberHttp()
	mongo := mongodb.NewMongoConnection()
	mongo.Info()

	healthCheck.HealthCheckRouter(http)

	err := http.ListenAndServe(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}

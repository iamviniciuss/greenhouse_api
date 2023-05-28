package infra

import (
	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
)

func TemperatureRouter(http infra.HttpService, temperatureRepository domain.TemperatureRepository) {
	http.Get("/greenhouse-api/v1/temperature/list", NewRegisterTemperatureCtrl(temperatureRepository).Execute)
}

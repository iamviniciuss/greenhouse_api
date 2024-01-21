package infra

import (
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	infra "github.com/iamviniciuss/greenhouse_api/src/infra/http"
)

func TemperatureRouter(http infra.HttpService, temperatureRepository repository.TemperatureRepository) {
	http.Get("/greenhouse-api/v1/temperature/list", NewRegisterTemperatureCtrl(temperatureRepository).Execute)
}

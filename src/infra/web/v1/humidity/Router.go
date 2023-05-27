package infra

import (
	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
)

func HumidityRouter(http infra.HttpService, humidityRepository domain.SoilRepository) {
	http.Get("/greenhouse-api/v1/humidity/list", NewRegisterHumidityCtrl(humidityRepository).List)
	http.Get("/greenhouse-api/v1/humidity", NewWaterPumpCtrl(humidityRepository).Execute)
	http.Post("/greenhouse-api/v1/humidity", NewRegisterHumidityCtrl(humidityRepository).Execute)
	http.Post("/greenhouse-api/v1/sensor", NewRegisterHumidityCtrl(humidityRepository).Sensor)
}

package infra

import (
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"

	infra "github.com/iamviniciuss/greenhouse_api/src/infra/http"
	list "github.com/iamviniciuss/greenhouse_api/src/infra/web/v1/humidity/list-all"
	manage_water_pump "github.com/iamviniciuss/greenhouse_api/src/infra/web/v1/humidity/manage-water-pump"
	register "github.com/iamviniciuss/greenhouse_api/src/infra/web/v1/humidity/register"
)

func HumidityRouter(http infra.HttpService, humidityRepository repository.SoilRepository) {
	http.Get("/greenhouse-api/v1/humidity/list", list.NewListAllHumidityCtrl(humidityRepository).List)
	http.Get("/greenhouse-api/v1/humidity", manage_water_pump.NewWaterPumpCtrl(humidityRepository).Execute)
	http.Post("/greenhouse-api/v1/humidity", register.NewRegisterHumidityCtrl(humidityRepository).Execute)
}

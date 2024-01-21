package infra

import (
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	infra "github.com/iamviniciuss/greenhouse_api/src/infra/http"
	create_sensor "github.com/iamviniciuss/greenhouse_api/src/infra/web/v1/sensor/create"
)

func SensorRouter(http infra.HttpService, humidityRepository repository.SoilRepository) {
	http.Post("/greenhouse-api/v1/sensor", create_sensor.NewCreateSensorCtrl(humidityRepository).Sensor)
}

package infra

import (
	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
	create_sensor "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/web/v1/sensor/create"
)

func SensorRouter(http infra.HttpService, humidityRepository domain.SoilRepository) {
	http.Post("/greenhouse-api/v1/sensor", create_sensor.NewCreateSensorCtrl(humidityRepository).Sensor)
}

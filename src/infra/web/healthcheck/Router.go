package infra

import (
	infra "github.com/iamviniciuss/greenhouse_api/src/infra/http"
	healthCheckCtrl "github.com/iamviniciuss/greenhouse_api/src/infra/web/healthcheck/controller"
)

func HealthCheckRouter(http infra.HttpService) {
	http.Get("/greenhouse-api/v1/health-check", healthCheckCtrl.HealthCheckCtrl)
	http.Get("/", healthCheckCtrl.HealthCheckCtrl)
}

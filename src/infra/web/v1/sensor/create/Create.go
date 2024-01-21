package infra

import (
	"encoding/json"

	"github.com/iamviniciuss/greenhouse_api/src/domain"
	infra "github.com/iamviniciuss/greenhouse_api/src/infra/errors"
	http "github.com/iamviniciuss/greenhouse_api/src/infra/http"
)

type SensorCtrlnput struct {
	Sensor domain.Sensor `json:"sensor"`
}

type CreateSensorCtrl struct {
	humidityRepository domain.SoilRepository
}

func NewCreateSensorCtrl(humidityRepository domain.SoilRepository) *CreateSensorCtrl {
	return &CreateSensorCtrl{
		humidityRepository: humidityRepository,
	}
}

func (wpc *CreateSensorCtrl) Sensor(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {

	var inputJSON SensorCtrlnput
	err := json.Unmarshal(body, &inputJSON)

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	created, err := wpc.humidityRepository.CreateSensor(&inputJSON.Sensor)

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	return created, nil
}

package infra

import (
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	infra "github.com/iamviniciuss/greenhouse_api/src/infra/errors"
	http "github.com/iamviniciuss/greenhouse_api/src/infra/http"
)

type RegisterTemperatureCtrlOutput struct {
	TurnOnWaterPump bool `json:"turn_on"`
}

type RegisterTemperatureCtrl struct {
	TemperatureRepository repository.TemperatureRepository
}

func NewRegisterTemperatureCtrl(temperatureRepository repository.TemperatureRepository) *RegisterTemperatureCtrl {
	return &RegisterTemperatureCtrl{
		TemperatureRepository: temperatureRepository,
	}
}

type RegisterTemperatureCtrlnput struct {
	Temperature float64 `json:"temperature,omitempty"`
	SensorID    string  `json:"sensor_id,omitempty"`
}

func (wpc *RegisterTemperatureCtrl) Execute(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {

	all, err := wpc.TemperatureRepository.ListAll()

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	return all, nil
}

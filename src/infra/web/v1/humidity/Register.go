package infra

import (
	"encoding/json"
	"fmt"

	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/errors"
	http "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
)

type RegisterHumidityCtrlOutput struct {
	TurnOnWaterPump bool `json:"turn_on"`
}

type RegisterHumidityCtrl struct {
	humidityRepository domain.TemperatureRepository
}

func NewRegisterHumidityCtrl(humidityRepository domain.TemperatureRepository) *RegisterHumidityCtrl {
	return &RegisterHumidityCtrl{
		humidityRepository: humidityRepository,
	}
}

type RegisterHumidityCtrlnput struct {
	Humidity int64  `json:"humidity,omitempty"`
	SensorID string `json:"sensor_id,omitempty"`
}

func (wpc *RegisterHumidityCtrl) Execute(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {

	var inputJSON RegisterHumidityCtrlnput
	err := json.Unmarshal(body, &inputJSON)

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	fmt.Println("Humidity: ", inputJSON.Humidity)
	created, err := wpc.humidityRepository.Create(&domain.HumidityRepositoryDTO{
		SensorID: inputJSON.SensorID,
		Value:    inputJSON.Humidity,
	})

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	return created, nil
}

type SensorCtrlnput struct {
	Sensor domain.Sensor `json:"sensor"`
}

func (wpc *RegisterHumidityCtrl) Sensor(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {

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

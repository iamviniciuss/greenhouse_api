package infra

import (
	"encoding/json"
	"fmt"

	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/errors"
	http "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
)

type RegisterTemperatureCtrlOutput struct {
	TurnOnWaterPump bool `json:"turn_on"`
}

type RegisterTemperatureCtrl struct {
	TemperatureRepository domain.TemperatureRepository
}

func NewRegisterTemperatureCtrl(temperatureRepository domain.TemperatureRepository) *RegisterTemperatureCtrl {
	return &RegisterTemperatureCtrl{
		TemperatureRepository: temperatureRepository,
	}
}

type RegisterTemperatureCtrlnput struct {
	Temperature int64  `json:"temperature,omitempty"`
	SensorID    string `json:"sensor_id,omitempty"`
}

func (wpc *RegisterTemperatureCtrl) Execute(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {

	var inputJSON RegisterTemperatureCtrlnput
	err := json.Unmarshal(body, &inputJSON)

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	fmt.Println("Temperature: ", inputJSON.Temperature)
	created, err := wpc.TemperatureRepository.Create(&domain.TemperatureRepositoryDTO{
		SensorID: inputJSON.SensorID,
		Value:    inputJSON.Temperature,
	})

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	return created, nil
}

type SensorCtrlnput struct {
	Sensor domain.Sensor `json:"sensor"`
}

func (wpc *RegisterTemperatureCtrl) Sensor(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {

	var inputJSON SensorCtrlnput
	err := json.Unmarshal(body, &inputJSON)

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	created, err := wpc.TemperatureRepository.CreateSensor(&inputJSON.Sensor)

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	return created, nil
}

func (wpc *RegisterTemperatureCtrl) List(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {

	all, err := wpc.TemperatureRepository.ListAll()

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	return all, nil
}

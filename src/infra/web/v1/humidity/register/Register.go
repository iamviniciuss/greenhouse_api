package infra

import (
	"encoding/json"
	"fmt"

	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/errors"
	http "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
)

type RegisterHumidityCtrl struct {
	humidityRepository domain.SoilRepository
}

func NewRegisterHumidityCtrl(humidityRepository domain.SoilRepository) *RegisterHumidityCtrl {
	return &RegisterHumidityCtrl{
		humidityRepository: humidityRepository,
	}
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

package infra

import (
	"encoding/json"
	"fmt"

	"github.com/iamviniciuss/greenhouse_api/src/application"
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	infra "github.com/iamviniciuss/greenhouse_api/src/infra/errors"
	http "github.com/iamviniciuss/greenhouse_api/src/infra/http"

	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
)

type RegisterHumidityCtrl struct {
	humidityRepository repository.SoilRepository
}

func NewRegisterHumidityCtrl(humidityRepository repository.SoilRepository) *RegisterHumidityCtrl {
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
	_, err = wpc.humidityRepository.Create(&repository.HumidityRepositoryDTO{
		SensorID: inputJSON.SensorID,
		Value:    inputJSON.Humidity,
	})

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	greenhouse := &shared.Greenhouse{
		ID:   "01",
		Name: "ESP32_HOUSE_VINICIUS",
		Sensors: []*shared.Sensor{
			{
				ID:            "645d82f4d2d163d2edc380a5",
				Envoironments: &shared.Envoironment{},
				Actuator: &shared.Actuator{
					ID:   "1",
					Name: "Bomba d'água",
				},
				IdealValue: []int{1400, 1450},
				Name:       "FC-28 - Sensor de umidade do solo",
			},
		},
	}

	command, err := application.NewManageWaterPump(wpc.humidityRepository).Execute(greenhouse)
	if err != nil {
		return nil, &infra.IntegrationError{
			StatusCode: 400,
			Message:    err.Error(),
		}
	}

	return command, nil
}

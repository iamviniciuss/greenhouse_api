package infra

import (
	application "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/application"
	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/errors"
	http "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
)

type WaterPumpCtrlOutput struct {
	TurnOnWaterPump bool `json:"turn_on"`
}

type WaterPumpCtrl struct {
	humidityRepository domain.SoilRepository
}

func NewWaterPumpCtrl(humidityRepository domain.SoilRepository) *WaterPumpCtrl {
	return &WaterPumpCtrl{
		humidityRepository: humidityRepository,
	}
}

func (wpc *WaterPumpCtrl) Execute(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {
	res, err := application.NewManageWaterPump(wpc.humidityRepository).Execute(&domain.Greenhouse{
		ID:   "01",
		Name: "ESP32_HOUSE_VINICIUS",
		Sensors: []*domain.Sensor{
			&domain.Sensor{
				ID:            "645d82f4d2d163d2edc380a5",
				Envoironments: &domain.Envoironment{},
				Actuator: &domain.Actuator{
					ID:   "1",
					Name: "Bomba d'água",
				},
				IdealValue: []int{800, 1638},
				Name:       "FC-28 - Sensor de umidade do solo",
			},
		},
	})

	if err != nil {
		return res, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	return &WaterPumpCtrlOutput{
		TurnOnWaterPump: res.TurnOnWaterPump,
	}, nil
}

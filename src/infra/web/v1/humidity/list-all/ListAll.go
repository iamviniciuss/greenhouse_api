package infra

import (
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/errors"
	http "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/http"
)

type ListAllHumidityCtrl struct {
	humidityRepository domain.SoilRepository
}

func NewListAllHumidityCtrl(humidityRepository domain.SoilRepository) *ListAllHumidityCtrl {
	return &ListAllHumidityCtrl{
		humidityRepository: humidityRepository,
	}
}

func (wpc *ListAllHumidityCtrl) List(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, *infra.IntegrationError) {

	all, err := wpc.humidityRepository.ListAll()

	if err != nil {
		return nil, &infra.IntegrationError{StatusCode: 400, Message: err.Error()}
	}

	return all, nil
}

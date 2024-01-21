package infra

import (
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	infra "github.com/iamviniciuss/greenhouse_api/src/infra/errors"
	http "github.com/iamviniciuss/greenhouse_api/src/infra/http"
)

type ListAllHumidityCtrl struct {
	humidityRepository repository.SoilRepository
}

func NewListAllHumidityCtrl(humidityRepository repository.SoilRepository) *ListAllHumidityCtrl {
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

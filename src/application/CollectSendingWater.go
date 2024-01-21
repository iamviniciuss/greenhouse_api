package application

import (
	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"
	soil "github.com/iamviniciuss/greenhouse_api/src/domain/soil"
)

type CollectSendingWater struct {
	repository.SoilRepository
}

type CollectSendingWaterOutput struct {
	EnergyConsume float64
	WaterBombed   float64
}

func NewCollectSendingWater(soilRepository repository.SoilRepository) *CollectSendingWater {
	return &CollectSendingWater{
		SoilRepository: soilRepository,
	}
}

func (m *CollectSendingWater) Execute(waterPump soil.WaterPump, event shared.Event) (CollectSendingWaterOutput, error) {
	energyConsumption := waterPump.EnergyConsumption(event.Duration()).EnergyConsumption
	producingWater := waterPump.ProducingWater(event.Duration()).NumberOfLitersPumped

	m.SoilRepository.RecordMetric(repository.MetricRepositoryDTO{
		Type:  shared.ENERGY_CONSUME,
		Value: energyConsumption,
	})

	m.SoilRepository.RecordMetric(repository.MetricRepositoryDTO{
		Type:  shared.WATER_BOMBED,
		Value: producingWater,
	})

	return CollectSendingWaterOutput{
		EnergyConsume: energyConsumption,
		WaterBombed:   producingWater,
	}, nil
}

package domain

type ProducingWaterOutput struct {
	NumberOfLitersPumped float64
}

func (wp *WaterPump) ProducingWater(numberOfMinutesConnected float64) ProducingWaterOutput {
	return ProducingWaterOutput{
		NumberOfLitersPumped: numberOfMinutesConnected * wp.ImpulseWaterPerMinutes,
	}
}

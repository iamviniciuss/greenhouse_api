package domain

type EnergyConsumptionOutput struct {
	EnergyConsumption float64
}

func (wp *WaterPump) CalculatePower() float64 {
	return wp.Voltage * wp.ElectricCurrent
}

func (wp *WaterPump) EnergyConsumption(numberOfMinutesConnected float64) EnergyConsumptionOutput {
	return EnergyConsumptionOutput{
		EnergyConsumption: wp.CalculatePower() * numberOfMinutesConnected,
	}
}

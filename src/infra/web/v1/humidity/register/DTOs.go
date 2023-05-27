package infra

type RegisterHumidityCtrlnput struct {
	Humidity int64  `json:"humidity,omitempty"`
	SensorID string `json:"sensor_id,omitempty"`
}

type RegisterHumidityCtrlOutput struct {
	TurnOnWaterPump bool `json:"turn_on"`
}

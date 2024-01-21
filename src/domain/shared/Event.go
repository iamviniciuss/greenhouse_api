package domain

import "time"

type Event struct {
	ID      string    `json:"_id" bson:"_id"`
	Started time.Time `json:"started" bson:"started"`
	Ended   time.Time `json:"ended" bson:"ended"`
}

func (eve *Event) Duration() float64 {
	return eve.Ended.Sub(eve.Started).Minutes()
}

type EventType string

const (
	ENERGY_CONSUME = "ENERGY_CONSUME"
	WATER_BOMBED   = "WATER_BOMBED"
)

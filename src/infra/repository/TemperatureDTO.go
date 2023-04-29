package repository

import "time"

type HumidityRepositoryDTO2 struct {
	ID          string    `json:"_id" bson:"_id"`
	Device      string    `json:"device" bson:"device"`
	CreatedAt   time.Time `json:"_created_at"`
	Temperature int64     `json:"temperature"`
}

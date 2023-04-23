package repository

import (
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/database"
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/database/mongodb"
)

type TemperatureRepositoryMongo[T mongodb.MongoInteface] struct {
	connection database.Connection[T]
}

func NewTemperatureRepositoryMongo(connection database.Connection[mongodb.MongoInteface]) *TemperatureRepositoryMongo[mongodb.MongoInteface] {
	return &TemperatureRepositoryMongo[mongodb.MongoInteface]{
		connection: connection,
	}
}

func (erm *TemperatureRepositoryMongo[T]) Create(temperature *Temperature) (*Temperature, error) {
	return &Temperature, nil
}

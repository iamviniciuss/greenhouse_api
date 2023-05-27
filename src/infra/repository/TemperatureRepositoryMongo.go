package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/database"
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/infra/database/mongodb"
	"github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/util/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo_lib "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TemperatureRepositoryMongo[T mongodb.MongoInteface] struct {
	connection database.Connection[T]
}

func NewTemperatureRepositoryMongo(connection database.Connection[mongodb.MongoInteface]) *TemperatureRepositoryMongo[mongodb.MongoInteface] {
	return &TemperatureRepositoryMongo[mongodb.MongoInteface]{
		connection: connection,
	}
}

func (erm *TemperatureRepositoryMongo[T]) CreateSensor(sensor *domain.Sensor) (*domain.Sensor, error) {
	var id primitive.ObjectID

	if sensor.ID == "" {
		id = primitive.NewObjectID()
	} else {
		id = mongo.GetObjectIDFromString(sensor.ID)
	}

	data := bson.M{
		"_id":           id,
		"name":          sensor.Name,
		"greenhouse_id": sensor.GreenhouseID,
		"ideal_value":   sensor.IdealValue,
	}

	res, err1 := erm.getCollection("sensor").InsertOne(context.TODO(), data)

	if err1 != nil {
		return nil, err1
	}

	sensor.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return sensor, nil
}

func (erm *TemperatureRepositoryMongo[T]) Create(humidity *domain.HumidityRepositoryDTO) (*domain.HumidityRepositoryDTO, error) {
	var id primitive.ObjectID

	if humidity.ID == "" {
		id = primitive.NewObjectID()
	} else {
		id = mongo.GetObjectIDFromString(humidity.ID)
	}

	readings := []float64{}
	last20Values, err := erm.FindLast20Values()
	if err != nil {
		return nil, err

	}

	for _, item := range last20Values {
		readings = append(readings, float64(item.Value))
	}

	humidity.CalculatePercentage()
	humidity.CalculateExponentialAverage(readings)
	humidity.CalculateMovelAverage(readings)

	data := bson.M{
		"_id":                 id,
		"created_at":          time.Now(),
		"sensor_id":           mongo.GetObjectIDFromString(humidity.SensorID),
		"value":               humidity.Value,
		"percentage":          humidity.Percentage,
		"exponential_average": humidity.ExponentialAverage,
		"movel_average":       humidity.MovelAverage,
	}

	res, err1 := erm.getCollection("humidity").InsertOne(context.TODO(), data)

	if err1 != nil {
		return nil, err1
	}

	humidity.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return humidity, nil
}

func (erm *TemperatureRepositoryMongo[T]) FindLast20Values() ([]*domain.HumidityRepositoryDTO, error) {
	findOptions := options.Find().SetSort(map[string]int{"created_at": -1})
	findOptions.SetLimit(20)
	result, err := erm.getCollection("humidity").Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	var data []*domain.HumidityRepositoryDTO
	err = result.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("not exists data")
	}

	return data, err
}

func (erm *TemperatureRepositoryMongo[T]) FindLastValue(temperature_id string) (*domain.HumidityRepositoryDTO, error) {
	findOptions := options.Find().SetSort(map[string]int{"created_at": -1})

	result, err := erm.getCollection("humidity").Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	var data []*domain.HumidityRepositoryDTO
	err = result.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("not exists data")
	}

	return data[0], err
}

func (erm *TemperatureRepositoryMongo[T]) FindSensorById(sensor_id string) (*domain.Sensor, error) {

	sensor := domain.Sensor{}
	err := erm.getCollection("sensor").
		FindOne(context.TODO(), bson.M{"_id": mongo.GetObjectIDFromString(sensor_id)}).
		Decode(&sensor)

	return &sensor, err
}

func (erm *TemperatureRepositoryMongo[T]) getCollection(collectionName string) *mongo_lib.Collection {
	return erm.connection.
		Client().
		Mongo().
		Database(os.Getenv("DATABASE")).
		Collection(collectionName)
}

func (erm *TemperatureRepositoryMongo[T]) ListAll() ([]*domain.HumidityRepositoryDTO, error) {
	findOptions := options.Find().SetSort(map[string]int{"created_at": -1})
	findOptions.SetLimit(1000)
	result, err := erm.getCollection("humidity").Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	var data []*domain.HumidityRepositoryDTO
	err = result.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("not exists data")
	}

	return data, err
}

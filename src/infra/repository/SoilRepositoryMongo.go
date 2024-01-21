package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	repository "github.com/iamviniciuss/greenhouse_api/src/domain/repository"
	shared "github.com/iamviniciuss/greenhouse_api/src/domain/shared"

	"github.com/iamviniciuss/greenhouse_api/src/infra/database"
	"github.com/iamviniciuss/greenhouse_api/src/infra/database/mongodb"
	"github.com/iamviniciuss/greenhouse_api/src/util/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo_lib "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SoilRepositoryMongo[T mongodb.MongoInteface] struct {
	connection database.Connection[T]
}

func NewSoilRepositoryMongo(connection database.Connection[mongodb.MongoInteface]) *SoilRepositoryMongo[mongodb.MongoInteface] {
	return &SoilRepositoryMongo[mongodb.MongoInteface]{
		connection: connection,
	}
}

func (erm *SoilRepositoryMongo[T]) CreateSensor(sensor *shared.Sensor) (*shared.Sensor, error) {
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

func (erm *SoilRepositoryMongo[T]) Create(humidity *repository.HumidityRepositoryDTO) (*repository.HumidityRepositoryDTO, error) {
	var id primitive.ObjectID

	if humidity.ID == "" {
		id = primitive.NewObjectID()
	} else {
		id = mongo.GetObjectIDFromString(humidity.ID)
	}

	humidity.CalculatePercentage()

	data := bson.M{
		"_id":        id,
		"created_at": time.Now(),
		"sensor_id":  mongo.GetObjectIDFromString(humidity.SensorID),
		"value":      humidity.Value,
		"percentage": humidity.Percentage,
	}

	res, err1 := erm.getCollection("humidity").InsertOne(context.TODO(), data)

	if err1 != nil {
		return nil, err1
	}

	humidity.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return humidity, nil
}

func (erm *SoilRepositoryMongo[T]) FindLast20Values() ([]*repository.HumidityRepositoryDTO, error) {
	findOptions := options.Find().SetSort(map[string]int{"created_at": -1})
	findOptions.SetLimit(20)
	result, err := erm.getCollection("humidity").Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	var data []*repository.HumidityRepositoryDTO
	err = result.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("not exists data")
	}

	return data, err
}

func (erm *SoilRepositoryMongo[T]) FindLastValue(humidity_id string) (*repository.HumidityRepositoryDTO, error) {
	findOptions := options.Find().SetSort(map[string]int{"created_at": -1})

	result, err := erm.getCollection("humidity").Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	var data []*repository.HumidityRepositoryDTO
	err = result.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("not exists data")
	}

	return data[0], err
}

func (erm *SoilRepositoryMongo[T]) FindSensorById(sensor_id string) (*shared.Sensor, error) {

	sensor := shared.Sensor{}
	err := erm.getCollection("sensor").
		FindOne(context.TODO(), bson.M{"_id": mongo.GetObjectIDFromString(sensor_id)}).
		Decode(&sensor)

	return &sensor, err
}

func (erm *SoilRepositoryMongo[T]) getCollection(collectionName string) *mongo_lib.Collection {
	return erm.connection.
		Client().
		Mongo().
		Database(os.Getenv("DATABASE")).
		Collection(collectionName)
}

func (erm *SoilRepositoryMongo[T]) ListAll() ([]*repository.HumidityRepositoryDTO, error) {
	findOptions := options.Find().SetSort(map[string]int{"created_at": -1})
	// findOptions.SetLimit(4000)
	result, err := erm.getCollection("humidity").Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	var data []*repository.HumidityRepositoryDTO
	err = result.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("not exists data")
	}

	return data, err
}

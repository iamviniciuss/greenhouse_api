package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoInteface interface {
	Mongo() *mongo.Client
}

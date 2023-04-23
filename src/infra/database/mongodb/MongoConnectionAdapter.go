package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConnection struct {
	client *mongo.Client
}

func NewMongoConnection() *MongoConnection {
	fmt.Println(os.Getenv("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		panic(err)
	}

	return &MongoConnection{
		client: client,
	}
}

func (mbc *MongoConnection) Info() {
	if err := mbc.client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Connection succeeded!")
}

func (mbc *MongoConnection) Client() MongoInteface {
	imp := NewMongoImp(mbc.client)
	return imp
}

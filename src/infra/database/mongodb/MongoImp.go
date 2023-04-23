package mongodb

import "go.mongodb.org/mongo-driver/mongo"

func NewMongoImp(cl *mongo.Client) *MongoImp {
	return &MongoImp{
		client: cl,
	}
}

type MongoImp struct {
	client *mongo.Client
}

func (mgi *MongoImp) Mongo() *mongo.Client {
	return mgi.client
}

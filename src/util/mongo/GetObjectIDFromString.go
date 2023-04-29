package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetObjectIDFromString(id string) primitive.ObjectID {

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return primitive.NilObjectID
	}

	return objectID
}

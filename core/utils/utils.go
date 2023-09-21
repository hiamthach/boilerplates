package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ObjectIDHex(s string) primitive.ObjectID {
	d, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return primitive.NilObjectID
	}
	return d
}

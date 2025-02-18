package mongotil

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func NewHexID() string {
	return NewID().Hex()
}

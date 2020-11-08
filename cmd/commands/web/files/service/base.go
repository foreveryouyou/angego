package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func toObjectID(s string) (objId primitive.ObjectID) {
	objId, _ = primitive.ObjectIDFromHex(s)
	return
}

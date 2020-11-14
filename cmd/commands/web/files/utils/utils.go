package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"strings"
)

func IsValidPhoneNumber(val string) bool {
	val = strings.TrimSpace(val)
	reg := regexp.MustCompile(`^1[0-9]{10}$`)
	return reg.MatchString(val)
}

func ToObjectID(s string) (objId primitive.ObjectID) {
	objId, _ = primitive.ObjectIDFromHex(s)
	return
}

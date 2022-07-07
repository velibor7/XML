package startup

import (
	"github.com/velibor7/XML/authentication_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var user_credentials = []*domain.UserCredential{
	{
		Id:       getObjectId("12306b1b644b3pa649s63l13"),
		Username: "velibor",
		Password: "velibor",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

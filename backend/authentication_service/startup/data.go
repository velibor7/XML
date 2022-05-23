package startup

import (
	"github.com/velibor7/XML/authentication_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var user_credentials = []*domain.UserCredential{
	{
		Id:       getObjectId("5789o4236423k5vjbhrwetkjbhbf980d"),
		Username: "velibor",
		Password: "test123",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

package startup

import (
	"github.com/velibor7/XML/authentication_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var user_credentials = []*domain.UserCredential{
	{
		Id:       getObjectId("48123asdv12j124bc834123asgh"),
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

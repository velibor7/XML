package startup

import (
	"github.com/velibor7/XML/auth_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:       primitive.NewObjectID(),
		Username: "velibor",
		Password: "velibor",
	},
	{
		Id:       primitive.NewObjectID(),
		Username: "Ln",
		Password: "Ln",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

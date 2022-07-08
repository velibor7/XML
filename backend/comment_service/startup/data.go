package startup

import (
	"github.com/velibor7/XML/comment_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var comments = []*domain.Comment{
	{
		Id:       getObjectId("63426daa12322d1748f63fe2"),
		Content:  "Svaka cast druze",
		Username: "velibor",
		PostId:   getObjectId("5789o4236423k5vjbhrwetkjbhbf980d"),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

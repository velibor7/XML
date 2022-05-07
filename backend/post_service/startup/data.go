package startup

import (
	"time"

	"github.com/velibor7/XML/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var post = []*domain.Post{
	{
		Id:       getObjectId("5789o4236423k5vjbhrwetkjbhbf980d"),
		Text:     "Neki post blablabla",
		Images:   "",
		Links:    "",
		Created:  time.Time{},
		Likes:    nil,
		Dislikes: nil,
		Comments: nil,
		UserId:   "1",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

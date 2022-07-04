package startup

import (
	"time"

	"github.com/velibor7/XML/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var posts = []*domain.Post{
	{
		Id:       getObjectId("5789o4236423k5vjbhrwetkjbhbf980d"),
		Text:     "Uspeli smo da napravimo i ovaj servis",
		Images:   "",
		Links:    "",
		Created:  time.Time{},
		Likes:    nil,
		Dislikes: nil,
		UserId:   getObjectId("48123asdv12j124bc834123asgh"),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

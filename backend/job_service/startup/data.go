package startup

import (
	"github.com/velibor7/XML/job_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jobs = []*domain.Job{
	{
		Id:          getObjectId("12306b12654b3pa649s63l13"),
		Title:       "developer",
		Description: "description",
		Skills:      nil,
		UserId:      getObjectId("12306b1b644b3pa649s63l13"),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

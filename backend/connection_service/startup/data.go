package startup

import (
	"time"

	"github.com/velibor7/XML/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var connections = []*domain.Connection{
	{
		Id:        getObjectId("64506b1d632b3sa762g63fe1"),
		IssuerId:  getObjectId("64506b1d632b3sa762g63fe2"),
		SubjectId: getObjectId("64506b1d632b3sa762g63fe4"),
		Approved:  true,
		Date:      time.Now(),
	},
}

var profilesPrivacy = []*domain.ProfilePrivacy{
	{
		Id:      primitive.NewObjectID(),
		UserId:  getObjectId("64506b1d632b3sa762g63fe2"),
		Private: false,
	},
	{
		Id:      primitive.NewObjectID(),
		UserId:  getObjectId("64506b1d632b3sa762g63fe4"),
		Private: false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

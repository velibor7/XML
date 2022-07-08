package startup

import (
	"time"

	"github.com/velibor7/XML/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var profiles = []*domain.Profile{
	{
		Id:             getObjectId("12306b1b644b3pa649s63l13"),
		Username:       "velibor",
		FirstName:      "Veljo",
		LastName:       "Vasic",
		FullName:       "VeliborVas",
		DateOfBirth:    time.Time{},
		PhoneNumber:    "012341123",
		Email:          "veljas7@gmail.com",
		Gender:         "male",
		Biography:      "Data Scientist",
		Education:      nil,
		WorkExperience: nil,
		Skills:         []string{"cooking", "labeling data"},
		Interests:      []string{"singing", "working out"},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

package api

import (
	"github.com/velibor7/XML/auth_service/domain"

	pb "github.com/velibor7/XML/common/proto/auth_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Function to return a pb.User from a domain.User
func mapUserToPb(user *domain.User) *pb.User {
	return &pb.User{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Password: user.Password,
	}
}

//Function to return a domain.User from a pb.User
func mapPbToUser(pbUser *pb.User) *domain.User {
	return &domain.User{
		Id:       getObjectId(pbUser.Id),
		Username: pbUser.Username,
		Password: pbUser.Password,
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

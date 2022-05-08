package api

import (
	"github.com/velibor7/XML/auth_service/domain"

	pb "github.com/velibor7/XML/common/proto/auth_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapUserToPb(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Password: user.Password,
	}
	return userPb
}

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

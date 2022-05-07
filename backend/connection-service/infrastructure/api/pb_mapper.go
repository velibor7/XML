package api

import (
	pb "github.com/velibor7/XML/tree/main/backend/common/proto/connection_service"
	"github.com/velibor7/XML/tree/main/backend/connection_service/domain"
)

func mapUserConnection(UserConnection *domain.UserConnection) *pb.User {
	UserConnectionPb := &pb.User{
		UserID:   UserConnection.UserID,
		IsPublic: UserConnection.IsPublic,
	}

	return UserConnectionPb
}

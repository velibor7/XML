package api

import (
	pb "github.com/velibor7/XML/common/proto/connection_service"
	"github.com/velibor7/XML/connection_service/domain"
)

func mapUserConn(userConn *domain.UserConn) *pb.User {
	userConnPb := &pb.User{
		UserID:   userConn.UserID,
		IsPublic: userConn.IsPublic,
	}

	return userConnPb
}

package domain

import pb "github.com/velibor7/XML/tree/main/backend/common/proto/connection_service"

type ConnectionStore interface {
	GetFriends(id string) ([]UserConnection, error)
	AddFriend(userIDa, userIDb string) (*pb.ActionResult, error)
	Register(userID string, isPublic bool) (*pb.ActionResult, error)
}

package domain

import pb "github.com/velibor7/XML/common/proto/connection_service"

type ConnectionStore interface {
	GetConnections(id string) ([]UserConn, error)
	AddConnection(userIDa, userIDb string) (*pb.ActionResult, error)
	Register(userID string, isPublic bool) (*pb.ActionResult, error)
	ApproveConnection(userIDa, userIDb string) (*pb.ActionResult, error)
	RejectConnection(userIDa, userIDb string) (*pb.ActionResult, error)
}

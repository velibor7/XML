package application

import (
	pb "github.com/velibor7/XML/common/proto/connection_service"
	"github.com/velibor7/XML/connection_service/domain"
)

type ConnectionService struct {
	store domain.ConnectionStore
}

func NewConnectionService(store domain.ConnectionStore) *ConnectionService {
	return &ConnectionService{
		store: store,
	}
}

func (service *ConnectionService) GetConnections(id string) ([]*domain.UserConn, error) {

	var friendsRetVal []*domain.UserConn

	friends, err := service.store.GetConnections(id)
	if err != nil {
		return nil, nil
	}
	for _, s := range friends {
		friendsRetVal = append(friendsRetVal, &domain.UserConn{UserID: s.UserID, IsPublic: s.IsPublic})
	}
	return friendsRetVal, nil
}

func (service *ConnectionService) Register(userID string, isPublic bool) (*pb.ActionResult, error) {
	return service.store.Register(userID, isPublic)
}

func (service *ConnectionService) AddConnection(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.AddConnection(userIDa, userIDb)
}

func (service *ConnectionService) ApproveConnection(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.ApproveConnection(userIDa, userIDb)
}

func (service *ConnectionService) RejectConnection(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.RejectConnection(userIDa, userIDb)
}

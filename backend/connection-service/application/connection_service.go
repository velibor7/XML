package application

import (
	pb "github.com/velibor7/XML/tree/main/backend/common/proto/connection_service"
	"github.com/velibor7/XML/tree/main/backend/connection_service/domain"
)

type ConnectionService struct {
	store domain.ConnectionStore
}

func NewConnectionService(store domain.ConnectionStore) *ConnectionService {
	return &ConnectionService{
		store: store,
	}
}

func (service *ConnectionService) GetFriends(id string) ([]*domain.UserConnection, error) {

	var friendsToReturn []*domain.UserConnection

	friends, err := service.store.GetFriends(id)
	if err != nil {
		return nil, err
	}
	for _, s := range friends {
		friendsToReturn = append(friendsToReturn, &domain.UserConnection{UserID: s.UserID, IsPublic: s.IsPublic})
	}
	return friendsToReturn, err
}

func (service *ConnectionService) Register(userID string, isPublic bool) (*pb.ActionResult, error) {
	return service.store.Register(userID, isPublic)
}

func (service *ConnectionService) AddFriend(userIDa, userIDb string) (*pb.ActionResult, error) {
	return service.store.AddFriend(userIDa, userIDb)
}

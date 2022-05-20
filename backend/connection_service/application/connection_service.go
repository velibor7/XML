package application

import (
	"github.com/velibor7/XML/connection_service/domain"
)

type ConnectionService struct {
	connInterface domain.ConnectionInterface
}

func NewConnectionService(inf domain.ConnectionInterface) *ConnectionService {
	return &ConnectionService{
		connInterface: inf,
	}
}

func (service *ConnectionService) Get(userId string) ([]*domain.Connection, error) {
	return service.connInterface.Get(userId)
}

func (service *ConnectionService) Create(conn *domain.Connection) (*domain.Connection, error) {
	return service.connInterface.Create(conn)
}

func (service *ConnectionService) Delete(id string) error {
	return service.connInterface.Delete(id)

}

func (service *ConnectionService) Update(id string) (*domain.Connection, error) {
	return service.connInterface.Update(id)
}

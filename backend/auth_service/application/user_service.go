package application

import "github.com/velibor7/XML/auth_service/domain"

type UserService struct {
	users domain.UserInterface
}

func NewUsersService(users domain.UserInterface) *UserService {
	return &UserService{
		users: users,
	}
}

func (service *UserService) Get(username string) (*domain.User, error) {
	return service.users.Get(username)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.users.GetAll()
}

func (service *UserService) Register(user *domain.User) error {
	return service.users.Register(user)
}

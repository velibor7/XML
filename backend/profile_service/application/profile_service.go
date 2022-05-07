package application

import (
	"github.com/velibor7/XML/profile_service/domain"
)

type ProfileService struct {
	profiles domain.ProfileInterface
}

func NewProfileService(profiles domain.ProfileInterface) *ProfileService {
	return &ProfileService{
		profiles: profiles,
	}
}

func (service *ProfileService) Get(username string) (*domain.Profile, error) {
	return service.profiles.Get(username)
}

func (service *ProfileService) GetAll(search string) ([]*domain.Profile, error) {
	return service.profiles.GetAll(search)
}

func (service *ProfileService) Create(profile *domain.Profile) error {
	return service.profiles.Create(profile)
}

func (service *ProfileService) Update(username string, profile *domain.Profile) error {
	return service.profiles.Update(username, profile)
}

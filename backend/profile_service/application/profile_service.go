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

func (service *ProfileService) Get(id string) (*domain.Profile, error) {
	return service.profiles.Get(id)
}

func (service *ProfileService) GetAll() ([]*domain.Profile, error) {
	return service.profiles.GetAll()
}

func (service *ProfileService) Create(profile *domain.Profile) error {
	return service.profiles.Create(profile)
}

func (service *ProfileService) Update(id string, profile *domain.Profile) error {
	return service.profiles.Update(id, profile)
}

///Rollback
func (service *ProfileService) RollbackUpdate(profile domain.Profile) error {
	return service.profiles.Update(profile.Id.Hex(), &profile)
}

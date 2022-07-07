package application

import (
	common "github.com/velibor7/XML/common/domain"
	"github.com/velibor7/XML/profile_service/domain"
)

type ProfileService struct {
	profiles     domain.ProfileInterface
	orchestrator *UpdateProfileOrchestrator
}

func NewProfileService(profiles domain.ProfileInterface, orchestrator *UpdateProfileOrchestrator) *ProfileService {
	return &ProfileService{
		profiles:     profiles,
		orchestrator: orchestrator,
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
	oldProfile, err := service.Get(id)

	err = service.profiles.Update(id, profile)
	if err != nil {
		return err
	}
	newProfile := &common.Profile{
		Id:             profile.Id,
		Username:       profile.Username,
		FirstName:      profile.FirstName,
		LastName:       profile.LastName,
		FullName:       profile.FirstName + profile.LastName,
		DateOfBirth:    profile.DateOfBirth,
		PhoneNumber:    profile.PhoneNumber,
		Email:          profile.Email,
		Gender:         profile.Gender,
		Biography:      profile.Biography,
		Education:      make([]common.Education, 0),
		WorkExperience: make([]common.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
		IsPrivate:      profile.IsPrivate,
	}
	for _, education := range profile.Education {
		education := &domain.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			Description:  education.Description,
		}
		profile.Education = append(profile.Education, *education)
	}

	for _, workExperience := range profile.WorkExperience {
		workExperience := &domain.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: workExperience.EmploymentType,
		}
		profile.WorkExperience = append(profile.WorkExperience, *workExperience)
	}

	for _, skill := range profile.Skills {
		profile.Skills = append(profile.Skills, skill)
	}

	for _, interest := range profile.Interests {
		profile.Interests = append(profile.Interests, interest)
	}
	err = service.orchestrator.Start(newProfile, oldProfile.Username, oldProfile.FirstName, oldProfile.LastName, oldProfile.IsPrivate)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProfileService) Delete(id string) error {
	return service.profiles.Delete(id)
}

func (service *ProfileService) RollbackUpdate(profile *domain.Profile) error {
	return service.profiles.Update(profile.Id.Hex(), profile)
}

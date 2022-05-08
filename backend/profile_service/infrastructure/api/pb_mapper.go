package api

import (
	pb "github.com/velibor7/XML/common/proto/profile_service"
	"github.com/velibor7/XML/profile_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapProfileToPb(profile *domain.Profile) *pb.Profile {
	pbProfile := &pb.Profile{
		Id:             profile.Id.Hex(),
		Username:       profile.Username,
		FirstName:      profile.FirstName,
		LastName:       profile.LastName,
		DateOfBirth:    timestamppb.New(profile.DateOfBirth),
		PhoneNumber:    profile.PhoneNumber,
		Email:          profile.Email,
		Gender:         profile.Gender,
		Biography:      profile.Biography,
		Education:      make([]*pb.Education, 0),
		WorkExperience: make([]*pb.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
	}

	for _, education := range profile.Education {
		educationPb := &pb.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			Description:  education.Description,
		}
		pbProfile.Education = append(pbProfile.Education, educationPb)
	}

	for _, workExperience := range profile.WorkExperience {
		workExperiencePb := &pb.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: 0,
		}
		pbProfile.WorkExperience = append(pbProfile.WorkExperience, workExperiencePb)
	}

	for _, skill := range profile.Skills {
		pbProfile.Skills = append(pbProfile.Skills, skill)
	}

	for _, interest := range profile.Interests {
		pbProfile.Interests = append(pbProfile.Interests, interest)
	}

	return pbProfile
}

func mapPbToProfile(pbProfile *pb.Profile) *domain.Profile {
	profile := &domain.Profile{
		Id:             getObjectId(pbProfile.Id),
		Username:       pbProfile.Username,
		FullName:       pbProfile.FullName,
		DateOfBirth:    pbProfile.DateOfBirth.AsTime(),
		PhoneNumber:    pbProfile.PhoneNumber,
		Email:          pbProfile.Email,
		Gender:         pbProfile.Gender,
		Biography:      pbProfile.Biography,
		Education:      make([]domain.Education, 0),
		WorkExperience: make([]domain.WorkExperience, 0),
		Skills:         make([]string, 0),
		Interests:      make([]string, 0),
	}
	for _, education := range pbProfile.Education {
		education := &domain.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,

			Description: education.Description,
		}
		profile.Education = append(profile.Education, *education)
	}

	for _, workExperience := range pbProfile.WorkExperience {
		workExperience := &domain.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: workExperience.EmploymentType.String(),
		}
		profile.WorkExperience = append(profile.WorkExperience, *workExperience)
	}

	for _, skill := range pbProfile.Skills {
		profile.Skills = append(profile.Skills, skill)
	}

	for _, interest := range pbProfile.Interests {
		profile.Interests = append(profile.Interests, interest)
	}
	return profile
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

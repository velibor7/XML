package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	Id             primitive.ObjectID `bson:"_id"`
	Username       string             `bson:"username"`
	FirstName      string             `bson:"firstName"`
	LastName       string             `bson:"lastName"`
	FullName       string             `bson:"fullName"`
	DateOfBirth    time.Time          `bson:"dateOfBirth"`
	PhoneNumber    string             `bson:"phoneNumber"`
	Email          string             `bson:"email"`
	Gender         string             `bson:"gender"`
	Biography      string             `bson:"biography"`
	Education      []Education        `bson:"education"`
	WorkExperience []WorkExperience   `bson:"workExperience"`
	Skills         []string           `bson:"skills"`
	Interests      []string           `bson:"interests"`
	IsPrivate      bool               `bson:"isPrivate"`
}

type Education struct {
	School       string `bson:"school"`
	Degree       string `bson:"degree"`
	FieldOfStudy string `bson:"fieldOfStudy"`
	Description  string `bson:"description"`
}

type WorkExperience struct {
	Title          string `bson:"title"`
	Company        string `bson:"company"`
	EmploymentType string `bson:"employmentType"`
}

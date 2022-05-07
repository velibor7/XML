package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	Id             primitive.ObjectID `bson:"_id"`
	Username       string             `bson:"username"`
	FullName       string             `bson:"fullName"`
	Email          string             `bson:"email"`
	Gender         string             `bson:"gender"`
	DateOfBirth    time.Time          `bson:"dateOfBirth"`
	PhoneNumber    string             `bson:"phoneNumber"`
	Biography      string             `bson:"biography"`
	Education      []Education        `bson:"education"`
	WorkExperience []WorkExperience   `bson:"workExperience"`
	Skills         []string           `bson:"skills"`
	Interests      []string           `bson:"interests"`
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

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
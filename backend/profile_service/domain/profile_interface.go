package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProfileInterface interface {
	Get(id string) (*Profile, error)
	GetAll() ([]*Profile, error)
	Create(profile *Profile) error
	Update(id primitive.ObjectID, profile *Profile) error
	DeleteAll() error
}

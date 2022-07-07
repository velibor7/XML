package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthStore interface {
	Login(user *UserCredential) (*UserCredential, error)
	Register(user *UserCredential) (*UserCredential, error)
	GetByUsername(username string) (*UserCredential, error)
	Delete(id primitive.ObjectID) error
	Update(id primitive.ObjectID, username string) (string, error)
	DeleteAll()
}

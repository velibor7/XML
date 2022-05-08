package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostInterface interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	Create(post *Post) (string, error)
	Update(post *Post) (string, error)
	// Delete(id primitive.ObjectID)
	DeleteAll()
	GetAllByUser(id string) ([]*Post, error)
}

package domain

type CommentInterface interface {
	GetForPost(id string) ([]*Comment, error)
	UpdateUsername(oldUsername, newUsername string) error
	Create(comment *Comment) error
	DeleteAll() error
}

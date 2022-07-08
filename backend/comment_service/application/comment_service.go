package application

import "github.com/velibor7/XML/comment_service/domain"

type CommentService struct {
	comments domain.CommentInterface
}

func NewCommentService(comments domain.CommentInterface) *CommentService {
	return &CommentService{
		comments: comments,
	}
}

func (service *CommentService) GetForPost(id string) ([]*domain.Comment, error) {
	return service.comments.GetForPost(id)
}

func (service *CommentService) UpdateUsername(oldUsername, newUsername string) error {
	return service.comments.UpdateUsername(oldUsername, newUsername)
}

func (service *CommentService) Create(comment *domain.Comment) error {
	return service.comments.Create(comment)
}

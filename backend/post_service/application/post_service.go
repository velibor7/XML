package application

import (
	"github.com/velibor7/XML/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	posts domain.PostInterface
}

func NewPostService(post_inf domain.PostInterface) *PostService {
	return &PostService{
		posts: post_inf,
	}
}

func (service *PostService) Get(id primitive.ObjectID) (*domain.Post, error) {
	return service.posts.Get(id)
}

func (service *PostService) GetAll() ([]*domain.Post, error) {
	return service.posts.GetAll()
}

func (service *PostService) Create(post *domain.Post) (string, error) {
	return service.posts.Create(post)
}

func (service *PostService) Update(post *domain.Post) (string, error) {
	return service.posts.Update(post)
}

// func (service *PostService) Delete(id primitive.ObjectID) error {
// 	return service.posts.Delete(id)
// }

func (service *PostService) GetAllByUser(id string) ([]*domain.Post, error) {
	return service.posts.GetAllByUser(id)
}

func (service *PostService) DeleteAll() {
	service.posts.DeleteAll()
}

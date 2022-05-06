package application

import (
	"github.com/velibor7/XML/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	posts domain.PostInterface
}

func NewPostService(service domain.PostInterface) *PostService {
	return &PostService{
		posts: posts,
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

func (service *PostService) Delete(post_id primitive.ObjectID) error {
	return service.posts.Delete(post_id)
}

func (service *PostService) DeleteAll() {
	service.posts.DeleteAll()
}

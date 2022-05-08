package api

import (
	"context"
	"fmt"

	pb "github.com/velibor7/XML/common/proto/post_service"
	"github.com/velibor7/XML/post_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	postPb := mapPost(post) // prepakujemo iz domenskog modela u protobuf oblik
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	posts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) GetAllByUser(ctx context.Context, request *pb.GetRequest) (*pb.GetAllResponse, error) {
	id := request.Id

	posts, err := handler.service.GetAllByUser(id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {

	// if request.Post.UserId == "" {
	// 	return nil, error(nil)
	// }

	fmt.Println(request.Post.Text)
	fmt.Println(request.Post.Id)
	post := mapCreatePost(request.Post)

	fmt.Println(post.Text)
	// userId := ctx.Value(interceptor.LoggedInUserKey{}).(string)
	// post.UserId = userId
	success, err := handler.service.Create(post)

	if err != nil {
		return nil, err
	}

	response := &pb.CreateResponse{
		Success: success,
	}

	return response, err
}

func (handler *PostHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Post.Id)
	oldPost, err := handler.service.Get(id)
	if err != nil {
		return &pb.UpdateResponse{
			Success: "error",
		}, err
	}

	post := mapUpdatePost(mapPost(oldPost), request.Post)
	success, err := handler.service.Update(post)
	response := &pb.UpdateResponse{
		Success: success,
	}
	return response, err
}

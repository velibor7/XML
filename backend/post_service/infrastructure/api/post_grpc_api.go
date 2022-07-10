package api

import (
	"context"
	"fmt"

	"github.com/velibor7/XML/common/loggers"
	pb "github.com/velibor7/XML/common/proto/post_service"
	"github.com/velibor7/XML/common/tracer"
	"github.com/velibor7/XML/post_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var log = loggers.NewPostLogger()

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
	span := tracer.StartSpanFromContext(ctx, "Get")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId)
	if err != nil {
		log.WithField("postId", request.Id).Errorf("Can't get post: %v", err)
		return nil, err
	}
	postPb := mapPost(post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	posts, err := handler.service.GetAll()
	if err != nil {
		log.Errorf("Can't get profile posts: %v", err)
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

func (handler *PostHandler) GetAllForUser(ctx context.Context, request *pb.GetRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAllForUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	id, err := primitive.ObjectIDFromHex(request.Id)
	posts, err := handler.service.GetAllForUser(id)
	if err != nil {
		log.WithField("profileId", request.Id).Errorf("Can't get profile posts: %v", err)
		return nil, err
	}
	fmt.Print(posts)
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
	span := tracer.StartSpanFromContext(ctx, "Create")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	post := mapCreatePost(request.Post)

	// userId := ctx.Value(interceptor.LoggedInUserKey{}).(string)
	// post.UserId = userId
	success, err := handler.service.Create(post)

	if err != nil {
		log.Errorf("Can't create post: %v", err)
		return nil, err
	}
	log.Info("Post created")
	response := &pb.CreateResponse{
		Success: success,
	}

	return response, nil
}

func (handler *PostHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Post.Id)
	oldPost, err := handler.service.Get(id)
	if err != nil {
		log.WithField("postId", id).Errorf("Can't get post: %v", err)
		return &pb.UpdateResponse{
			Success: "error",
		}, err
	}

	post := mapUpdatePost(mapPost(oldPost), request.Post)
	success, err := handler.service.Update(post)
	if err != nil {
		log.WithField("postId", id).Errorf("Can't update post: %v", err)
		return &pb.UpdateResponse{
			Success: "error",
		}, err
	}
	log.WithField("postId", id).Infof("Post updated")
	response := &pb.UpdateResponse{
		Success: success,
	}
	return response, nil
}

package api

import (
	"context"

	"github.com/velibor7/XML/comment_service/application"
	"github.com/velibor7/XML/common/loggers"
	pb "github.com/velibor7/XML/common/proto/comment_service"
)

var log = loggers.NewCommentLogger()

type CommentHandler struct {
	pb.UnimplementedCommentServiceServer
	service *application.CommentService
}

func NewCommentHandler(service *application.CommentService) *CommentHandler {
	return &CommentHandler{
		service: service,
	}
}

func (handler *CommentHandler) GetForPost(ctx context.Context, request *pb.GetForPostRequest) (*pb.GetForPostResponse, error) {
	comments, err := handler.service.GetForPost(request.Id)
	if err != nil {
		log.WithField("postId", request.Id).Errorf("Cannot get comments: %v", err)
		return nil, err
	}
	response := &pb.GetForPostResponse{
		Comments: []*pb.Comment{},
	}

	for _, comment := range comments {
		temp := mapComment(comment)
		response.Comments = append(response.Comments, temp)
	}
	return response, nil
}

func (handler *CommentHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	comment := mapPb(request.Comment)
	err := handler.service.Create(comment)
	if err != nil {
		log.Errorf("Cannot create comment: %v", err)
		return nil, err
	}
	log.Info("Comment created: %v", comment.Id)
	return &pb.CreateResponse{
		Message: "success",
	}, nil
}

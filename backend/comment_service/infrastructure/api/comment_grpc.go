package api

import (
	"context"

	"github.com/velibor7/XML/comment_service/application"
	pb "github.com/velibor7/XML/common/proto/comment_service"
)

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
	comments := handler.service.GetForPost(request.Id)
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
		return nil, err
	}

	return &pb.CreateResponse{
		Message: "success",
	}, nil
}

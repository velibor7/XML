package api

import (
	"context"

	"github.com/velibor7/XML/auth_service/application"
	pb "github.com/velibor7/XML/common/proto/auth_service"
)

type UserHandler struct {
	pb.UnimplementedAuthenticationServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	username := request.Username
	objectUsername := username

	product, err := handler.service.Get(objectUsername)
	if err != nil {
		return nil, err
	}
	productPb := mapUserToPb(product)
	response := &pb.GetResponse{
		User: productPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, User := range Users {
		current := mapUserToPb(User)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := mapPbToUser(request.User)
	err := handler.service.Register(user)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		User: mapUserToPb(user),
	}, nil

}

package api

import (
	"context"

	"github.com/velibor7/XML/authentication_service/application"
	pb "github.com/velibor7/XML/common/proto/authentication_service"
	pbProfile "github.com/velibor7/XML/common/proto/profile_service"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service       *application.AuthService
	profileClient pbProfile.ProfileServiceClient
}

func NewAuthHandler(service *application.AuthService, profileClient pbProfile.ProfileServiceClient) *AuthHandler {
	return &AuthHandler{
		service:       service,
		profileClient: profileClient,
	}
}

func (handler *AuthHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.JWTResponse, error) {
	userCredential := mapPbUserCredential(request.User)
	jwt, err := handler.service.Login(userCredential)
	if err != nil {
		return nil, err
	}
	return &pb.JWTResponse{
		Token:    jwt.Token,
		UserId:   jwt.UserId.Hex(),
		Username: jwt.Username,
	}, nil
}

func (handler *AuthHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	userCredential := mapPbUserCredential(request.User)
	user, err := handler.service.Register(userCredential)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		User: mapUserCredential(user),
	}, nil
}

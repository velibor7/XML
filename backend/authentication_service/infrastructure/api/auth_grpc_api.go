package api

import (
	"context"

	"github.com/velibor7/XML/authentication_service/application"
	"github.com/velibor7/XML/common/loggers"
	pb "github.com/velibor7/XML/common/proto/authentication_service"
	pbProfile "github.com/velibor7/XML/common/proto/profile_service"
)

var log = loggers.NewAuthenticationLogger()

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
		log.WithField("username", jwt.Username).Errorf("Log in app : %v", err)
		return nil, err
	}
	log.Info("Login!")
	return &pb.JWTResponse{
		Token:    jwt.Token,
		Id:       jwt.Id.Hex(),
		Username: jwt.Username,
	}, nil
}

func (handler *AuthHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	userCredential := mapPbUserCredential(request.User)
	user, err := handler.service.Register(userCredential)
	if err != nil {
		log.Errorf("Registration falied: %v", err)
		return nil, err
	}
	log.Info("Registered!")
	return &pb.RegisterResponse{
		User: mapUserCredential(user),
	}, nil
}

package application

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/velibor7/XML/authentication_service/domain"
	auth "github.com/velibor7/XML/common/domain"
	"github.com/velibor7/XML/common/loggers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var log = loggers.NewAuthenticationLogger()

type AuthService struct {
	store        domain.AuthStore
	orchestrator *CreateProfileOrchestrator
}

func NewAuthService(store domain.AuthStore, orchestrator *CreateProfileOrchestrator) *AuthService {
	return &AuthService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *AuthService) Login(user *domain.UserCredential) (*domain.JWTTokenFullResponse, error) {
	user, err := service.store.Login(user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "User was not found")
	}
	token, err := service.GenerateJWT(user.Username, user.Id.Hex())
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (service *AuthService) Register(user *domain.UserCredential) (*domain.UserCredential, error) {
	regUser, err := service.store.Register(user)
	if err != nil {
		return nil, err
	}
	fmt.Print("----regUser----\n")
	fmt.Print(regUser)
	profile := &auth.Profile{
		Id:             regUser.Id,
		Username:       regUser.Username,
		FirstName:      "",
		LastName:       "",
		FullName:       "",
		DateOfBirth:    time.Time{},
		PhoneNumber:    "",
		Email:          "",
		Gender:         "",
		IsPrivate:      false,
		Biography:      "",
		Education:      nil,
		WorkExperience: nil,
		Skills:         nil,
		Interests:      nil,
	}
	err = service.orchestrator.Start(profile)
	if err != nil {
		log.Errorf("Saga failed: %v", err)
		service.store.Delete(user.Id)
		return nil, err
	}
	return user, nil

}

func (service *AuthService) GenerateJWT(username, id string) (*domain.JWTTokenFullResponse, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("Something Went Wrong: %s", err.Error())
		return nil, err
	}
	Id, _ := primitive.ObjectIDFromHex(id)
	return &domain.JWTTokenFullResponse{
		Token:    tokenString,
		Username: username,
		Id:       Id,
	}, nil
}

func (service *AuthService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *AuthService) Update(id primitive.ObjectID, username string) (string, error) {
	return service.store.Update(id, username)
}

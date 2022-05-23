package application

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/velibor7/XML/authentication_service/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	store domain.AuthStore
}

func NewAuthService(store domain.AuthStore) *AuthService {
	return &AuthService{
		store: store,
	}
}

func (service *AuthService) Login(user *domain.UserCredential) (*domain.JWTToken, error) {
	user, err := service.store.Login(user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "User was not found")
	}
	token, err := service.GenerateJWT(user.Username)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (service *AuthService) Register(user *domain.UserCredential) (*domain.UserCredential, error) {
	user, err := service.store.Register(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *AuthService) GenerateJWT(username string) (*domain.JWTToken, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("Something Went Wrong: %s", err.Error())
		return nil, err
	}
	return &domain.JWTToken{Token: tokenString}, nil
}

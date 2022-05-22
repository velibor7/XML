package application

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/velibor7/XML/authentication_service/domain"
)

type JWTService struct {
	privateKey          *rsa.PrivateKey
	publicKey           *rsa.PublicKey
	accessTokenDuration time.Duration
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func NewJWTManager(privateKey, publicKey string) (*JWTService, error) {
	parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	parsedPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil, err
	}
	return &JWTService{
		privateKey:          parsedPrivateKey,
		publicKey:           parsedPublicKey,
		accessTokenDuration: 10 * time.Minute, // TODO: bilo je 10
	}, nil
}

func (manager *JWTService) GenerateToken(auth *domain.Authentication) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   auth.Id,
			ExpiresAt: time.Now().Add(manager.accessTokenDuration).Unix(),
		},
		Username: auth.Username,
		Role:     auth.Role,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		claims,
	)
	return token.SignedString(manager.privateKey)
}

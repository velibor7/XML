package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/velibor7/XML/api_gateway/domain"
	"github.com/velibor7/XML/api_gateway/infrastructure/services"
	authentication "github.com/velibor7/XML/common/proto/auth_service"
)

type AuthHandler struct {
	authClientAddress string
}

func newAuthHandler(authClientAddress string) Handler {
	return &AuthHandler{
		authClientAddress: authClientAddress,
	}

}

func (handler *AuthHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/{username}", handler.GetUserInfo)
	if err != nil {
		panic(err)
	}
}

func (hadler *AuthHandler) GetUserInfo(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	username := pathParams["username"]
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := &domain.User{Username: "username"}

	err := hadler.addUserInfo(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *AuthHandler) addUserInfo(user *domain.User) error {
	authClient := services.NewAuthenticationClient(handler.authClientAddress)
	userInfo, err := authClient.Get(context.TODO(), &authentication.GetRequest{Username: user.Username})
	if err != nil {
		panic(err)
	}
	user.Username = userInfo.User.Username
	user.Id = userInfo.User.Id
	user.Password = userInfo.User.Password

	return nil
}

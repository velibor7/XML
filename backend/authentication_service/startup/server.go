package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/velibor7/XML/authentication_service/application"
	"github.com/velibor7/XML/authentication_service/domain"
	"github.com/velibor7/XML/authentication_service/infrastructure/api"
	"github.com/velibor7/XML/authentication_service/infrastructure/persistence"
	"github.com/velibor7/XML/authentication_service/startup/config"
	auth "github.com/velibor7/XML/common/proto/authentication_service"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	mongoClient := server.initMongoClient()
	authStore := server.initAuthStore(mongoClient)
	authService := server.initAuthService(authStore)
	authHandler := server.initAuthHandler(authService)
	server.startGrpcServer(authHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthDBHost, server.config.AuthDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAuthStore(client *mongo.Client) domain.AuthStore {
	store := persistence.NewAuthMongoDBStore(client)
	store.DeleteAll()
	for _, UserCredential := range user_credentials {
		_, err := store.Register(UserCredential)
		if err != nil {
			panic(err)
		}

	}
	return store
}

func (server *Server) initAuthService(store domain.AuthStore) *application.AuthService {
	return application.NewAuthService(store)
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.AuthHandler {
	return api.NewAuthHandler(service)
}

func (server *Server) startGrpcServer(authHandler *api.AuthHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

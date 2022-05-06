package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/velibor7/XML/auth_service/application"
	"github.com/velibor7/XML/auth_service/domain"
	"github.com/velibor7/XML/auth_service/infrastructure/api"
	"github.com/velibor7/XML/auth_service/infrastructure/persistence"
	"github.com/velibor7/XML/auth_service/startup/config"
	security "github.com/velibor7/XML/common/proto/auth_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	//
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "auth_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userInterface := server.initUserInterface(mongoClient)
	authenticationService := server.initAuthenticationService(userInterface)
	userHandler := server.initUserHandler(authenticationService)
	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthenticationDBHost, server.config.AuthenticationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserInterface(client *mongo.Client) domain.UserInterface {
	inf := persistence.NewUserMongoDBInterface(client)
	err := inf.DeleteAll()

	if err != nil {
		return nil
	}

	for _, User := range users {
		err := inf.Register(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	return inf
}

func (server *Server) initAuthenticationService(inf domain.UserInterface) *application.UserService {
	return application.NewUsersService(inf)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	security.RegisterAuthenticationServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

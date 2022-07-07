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
	client "github.com/velibor7/XML/common/client"
	auth "github.com/velibor7/XML/common/proto/authentication_service"
	profile "github.com/velibor7/XML/common/proto/profile_service"
	saga "github.com/velibor7/XML/common/saga/messaging"
	"github.com/velibor7/XML/common/saga/messaging/nats"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

const (
	QueueGroup = "auth_service"
)

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	mongoClient := server.initMongoClient()
	authStore := server.initAuthStore(mongoClient)

	commandPublisher := server.initPublisher(server.config.CreateProfileCommandSubject)
	replySubscriber := server.initSubscriber(server.config.CreateProfileReplySubject, QueueGroup)
	createProfileOrchestrator := server.initCreateProfileOrchestrator(commandPublisher, replySubscriber)

	authService := server.initAuthService(authStore, createProfileOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.CreateProfileCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateProfileReplySubject)
	server.initCreateProfileHandler(authService, replyPublisher, commandSubscriber)

	commandSubscriber = server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	replyPublisher = server.initPublisher(server.config.UpdateProfileReplySubject)
	server.initUpdateProfileHandler(authService, replyPublisher, commandSubscriber)

	profileClient, err := client.NewProfileClient(fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort))
	if err != nil {
		log.Fatalf("PCF: %v", err)
	}

	authHandler := server.initAuthHandler(authService, profileClient)
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

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initCreateProfileOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.CreateProfileOrchestrator {
	orchestrator, err := application.NewCreateProfileOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}
func (server *Server) initAuthService(store domain.AuthStore, createProfileOrchestrator *application.CreateProfileOrchestrator) *application.AuthService {
	return application.NewAuthService(store, createProfileOrchestrator)
}

func (server *Server) initCreateProfileHandler(service *application.AuthService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUpdateProfileHandler(service *application.AuthService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initAuthHandler(service *application.AuthService, profileClient profile.ProfileServiceClient) *api.AuthHandler {
	return api.NewAuthHandler(service, profileClient)
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

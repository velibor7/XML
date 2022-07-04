package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/velibor7/XML/common/client"
	pbComment "github.com/velibor7/XML/common/proto/comment_service"
	profile "github.com/velibor7/XML/common/proto/profile_service"
	saga "github.com/velibor7/XML/common/saga/messaging"
	"github.com/velibor7/XML/common/saga/messaging/nats"
	"github.com/velibor7/XML/profile_service/application"
	"github.com/velibor7/XML/profile_service/domain"
	"github.com/velibor7/XML/profile_service/infrastructure/api"
	"github.com/velibor7/XML/profile_service/infrastructure/persistence"
	"github.com/velibor7/XML/profile_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	QueueGroup = "profile_service"
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
	profileInterface := server.initProfileInterface(mongoClient)

	commandPublisher := server.initPublisher(server.config.UpdateProfileCommandSubject)
	replySubscriber := server.initSubscriber(server.config.UpdateProfileReplySubject, QueueGroup)
	updateProfileOrchestrator := server.initUpdateProfileOrchestrator(commandPublisher, replySubscriber)

	profileService := server.initProfileService(profileInterface, updateProfileOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.UpdateProfileReplySubject)
	server.initUpdateProfileHandler(profileService, replyPublisher, commandSubscriber)

	commentClient, err := client.NewCommentClient(fmt.Sprintf("%s:%s", server.config.CommentHost, server.config.CommentPort))
	if err != nil {
		log.Fatal(err)
	}

	profileHandler := server.initProfileHandler(profileService, commentClient)
	server.startGrpcServer(profileHandler)

}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ProfileDBHost, server.config.ProfileDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initProfileInterface(client *mongo.Client) domain.ProfileInterface {
	inf := persistence.NewProfileMongoDB(client)
	err := inf.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, Profile := range profiles {
		err := inf.Create(Profile)
		if err != nil {
			panic(err)
		}

	}
	return inf
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

func (server *Server) initUpdateProfileOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.UpdateProfileOrchestrator {
	orchestrator, err := application.NewUpdateProfileOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}
func (server *Server) initProfileService(inf domain.ProfileInterface, orchestrator *application.UpdateProfileOrchestrator) *application.ProfileService {
	return application.NewProfileService(inf, orchestrator)
}
func (server *Server) initUpdateProfileHandler(service *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initProfileHandler(service *application.ProfileService,
	commentClient pbComment.CommentServiceClient) *api.ProfileHandler {
	return api.NewProfileHandler(service, commentClient)
}

func (server *Server) startGrpcServer(profileHandler *api.ProfileHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	profile.RegisterProfileServiceServer(grpcServer, profileHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/velibor7/XML/comment_service/application"
	"github.com/velibor7/XML/comment_service/domain"
	"github.com/velibor7/XML/comment_service/infrastructure/api"
	"github.com/velibor7/XML/comment_service/infrastructure/persistence"
	"github.com/velibor7/XML/comment_service/startup/config"
	comment "github.com/velibor7/XML/common/proto/comment_service"
	saga "github.com/velibor7/XML/common/saga/messaging"
	"github.com/velibor7/XML/common/saga/messaging/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

const (
	QueueGroup = "comment_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	fmt.Println("klijent je pokrenut")
	commentInterface := server.initCommentInterface(mongoClient)
	fmt.Println("interface je pokrenut")
	commentService := server.initCommentService(commentInterface)
	fmt.Println("service je pokrenut")
	commandSubscriber := server.initSubscriber(server.config.UpdateProfileCommandSubject, QueueGroup)
	fmt.Println("sub je pokrenut")
	replyPublisher := server.initPublisher(server.config.UpdateProfileReplySubject)
	fmt.Println("pub je pokrenut")
	server.initUpdateProfileHandler(commentService, replyPublisher, commandSubscriber)
	fmt.Println("update je pokrenut")
	commentHandler := server.initCommentHandler(commentService)
	fmt.Println("handler je pokrenut")
	server.startGrpcServer(commentHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.CommentDBHost, server.config.CommentDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initCommentInterface(client *mongo.Client) domain.CommentInterface {
	inf := persistence.NewCommentMongoDB(client)
	err := inf.DeleteAll()
	if err != nil {
		return nil
	}
	for _, comment := range comments {
		err := inf.Create(comment)
		if err != nil {
			log.Fatal(err)
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

func (server *Server) initCommentService(inf domain.CommentInterface) *application.CommentService {
	return application.NewCommentService(inf)
}

func (server *Server) initUpdateProfileHandler(service *application.CommentService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateProfileCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initCommentHandler(service *application.CommentService) *api.CommentHandler {
	return api.NewCommentHandler(service)
}

func (server *Server) startGrpcServer(commentHandler *api.CommentHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	comment.RegisterCommentServiceServer(grpcServer, commentHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

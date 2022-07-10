package startup

import (
	"fmt"
	"io"
	"net"

	otgo "github.com/opentracing/opentracing-go"
	"github.com/velibor7/XML/common/loggers"
	post "github.com/velibor7/XML/common/proto/post_service"
	"github.com/velibor7/XML/common/tracer"
	"github.com/velibor7/XML/post_service/application"
	"github.com/velibor7/XML/post_service/domain"
	"github.com/velibor7/XML/post_service/infrastructure/api"
	"github.com/velibor7/XML/post_service/infrastructure/persistence"
	"github.com/velibor7/XML/post_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log = loggers.NewPostLogger()

type Server struct {
	config *config.Config
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init("post_service")
	return &Server{
		tracer: tracer,
		config: config,
		closer: closer,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	postInterface := server.initPostInterface(mongoClient)
	postService := server.initPostService(postInterface)
	postHandler := server.initPostHandler(postService)

	server.startGrpcServer(postHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initPostInterface(client *mongo.Client) domain.PostInterface {
	inf := persistence.NewPostMongoDB(client)
	inf.DeleteAll()
	for _, Post := range posts {
		_, err := inf.Create(Post)
		if err != nil {
			log.Fatal(err)
		}

	}
	return inf
}

func (server *Server) initPostService(inf domain.PostInterface) *application.PostService {
	return application.NewPostService(inf)
}

func (server *Server) initPostHandler(service *application.PostService) *api.PostHandler {
	return api.NewPostHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.PostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	post.RegisterPostServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

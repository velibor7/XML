package startup

import (
	"fmt"
	"log"
	"net"

	job "github.com/velibor7/XML/common/proto/job_service"
	"github.com/velibor7/XML/job_service/application"
	"github.com/velibor7/XML/job_service/domain"
	"github.com/velibor7/XML/job_service/infrastructure/api"
	"github.com/velibor7/XML/job_service/infrastructure/persistence"
	"github.com/velibor7/XML/job_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	jobInterface := server.initJobInterface(mongoClient)

	jobService := server.initJobService(jobInterface)
	jobHandler := server.initJobHandler(jobService)

	server.startGrpcServer(jobHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.JobDBHost, server.config.JobDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initJobInterface(client *mongo.Client) domain.JobInterface {
	jobInterface := persistence.NewJobMongoDB(client)
	err := jobInterface.DeleteAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, Job := range jobs {
		err := jobInterface.Create(Job)
		if err != nil {
			panic(err)
		}
	}
	return jobInterface
}

func (server *Server) initJobService(inf domain.JobInterface) *application.JobService {
	return application.NewJobService(inf)
}

func (server *Server) initJobHandler(service *application.JobService) *api.JobHandler {
	return api.NewJobHandler(service)
}

func (server *Server) startGrpcServer(jobHandler *api.JobHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	job.RegisterJobServiceServer(grpcServer, jobHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

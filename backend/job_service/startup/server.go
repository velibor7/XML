package startup

import (
	"fmt"
	"net"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/velibor7/XML/common/loggers"
	job "github.com/velibor7/XML/common/proto/job_service"
	"github.com/velibor7/XML/job_service/application"
	"github.com/velibor7/XML/job_service/domain"
	"github.com/velibor7/XML/job_service/infrastructure/api"
	"github.com/velibor7/XML/job_service/infrastructure/persistence"
	"github.com/velibor7/XML/job_service/startup/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log = loggers.NewJobLogger()

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	neo4jClient := server.initNeo4J()

	jobInterface := server.initJobInterface(neo4jClient)

	jobService := server.initJobService(jobInterface)
	jobHandler := server.initJobHandler(jobService)

	server.startGrpcServer(jobHandler)
}

func (server *Server) initNeo4J() *neo4j.Driver {
	uri := "neo4j://neo4j:7687"
	client, err := persistence.GetClient(uri, server.config.Neo4jUsername, server.config.Neo4jPassword)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initJobInterface(client *neo4j.Driver) domain.JobInterface {
	return persistence.NewJobNeo4j(client)

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

package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	inventory "github.com/velibor7/XML/tree/main/backend/common/proto/connection_service"
	"github.com/velibor7/XML/tree/main/backend/connection_service/application"
	"github.com/velibor7/XML/tree/main/backend/connection_service/domain"
	"github.com/velibor7/XML/tree/main/backend/connection_service/infrastructure/api"
	"github.com/velibor7/XML/tree/main/backend/connection_service/infrastructure/persistence"
	"github.com/velibor7/XML/tree/main/backend/connection_service/startup/config"
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

const (
	QueueGroup = "connection_service"
)

func (server *Server) Start() {
	neo4jClient := server.initNeo4J()

	connectionStore := server.initConnectionStore(neo4jClient)

	connectionService := server.initConnectionService(connectionStore)

	connectionHandler := server.initConnectionHandler(connectionService)

	server.startGrpcServer(connectionHandler)
}

func (server *Server) initNeo4J() *neo4j.Driver {
	dbUri := "bolt://localhost:7687"

	client, err := persistence.GetClient(dbUri, server.config.Neo4jUsername, server.config.Neo4jPassword)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initConnectionStore(client *neo4j.Driver) domain.ConnectionStore {
	store := persistence.NewConnectionDBStore(client)
	return store
}

func (server *Server) initConnectionService(store domain.ConnectionStore) *application.ConnectionService {
	return application.NewConnectionService(store)
}

func (server *Server) initConnectionHandler(service *application.ConnectionService) *api.ConnectionHandler {
	return api.NewConnectionHandler(service)
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	inventory.RegisterConnectionServiceServer(grpcServer, connectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

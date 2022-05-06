package startup

import (
	"log"

	"github.com/velibor7/XML/profile_service/infrastructure/persistence"
	"github.com/velibor7/XML/profile_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func Start(server *Server) {
	//mongoClient := server.initMongoClient()
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ProfileDBHost, server.config.ProfileDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

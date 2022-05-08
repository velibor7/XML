package startup

import (
	"fmt"
	"log"
	"net"

	profile "github.com/velibor7/XML/common/proto/profile_service"
	"github.com/velibor7/XML/profile_service/application"
	"github.com/velibor7/XML/profile_service/domain"
	"github.com/velibor7/XML/profile_service/infrastructure/api"
	"github.com/velibor7/XML/profile_service/infrastructure/persistence"
	"github.com/velibor7/XML/profile_service/startup/config"
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
	profileInterface := server.initProfileInterface(mongoClient)
	profileService := server.initProfileService(profileInterface)
	profileHandler := server.initProfileHandler(profileService)

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

func (server *Server) initProfileService(inf domain.ProfileInterface) *application.ProfileService {
	return application.NewProfileService(inf)
}

func (server *Server) initProfileHandler(service *application.ProfileService) *api.ProfileHandler {
	return api.NewProfileHandler(service)
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

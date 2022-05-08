package startup

import (
	"fmt"
	"log"
	"net"
	"time"

	profile "github.com/velibor7/XML/common/proto/profile_service"
	"github.com/velibor7/XML/profile_service/application"
	"github.com/velibor7/XML/profile_service/domain"
	"github.com/velibor7/XML/profile_service/infrastructure/api"
	"github.com/velibor7/XML/profile_service/infrastructure/persistence"
	"github.com/velibor7/XML/profile_service/startup/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Start(server *Server) {
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

func (server *Server) startGrpcServer(userHandler *api.ProfileHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	profile.RegisterProfileServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

var profiles = []*domain.Profile{
	{
		Id:             getObjectId("12306b1b644b3pa649s63l13"),
		Username:       "velibor",
		FullName:       "VeliborVas",
		DateOfBirth:    time.Time{},
		PhoneNumber:    "012341123",
		Email:          "veljas7@gmail.com",
		Gender:         "male",
		Biography:      "Data Scientist",
		Education:      nil,
		WorkExperience: nil,
		Skills:         nil,
		Interests:      nil,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

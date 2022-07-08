package api

import (
	"context"

	"github.com/velibor7/XML/common/loggers"
	pb "github.com/velibor7/XML/common/proto/connection_service"
	"github.com/velibor7/XML/connection_service/application"
)

var log = loggers.NewConnectionLogger()

type ConnectionHandler struct {
	pb.UnimplementedConnectionServiceServer
	service *application.ConnectionService
}

func NewConnectionHandler(service *application.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{
		service: service,
	}
}

func (handler *ConnectionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	connections, err := handler.service.Get(request.UserId)
	if err != nil {
		log.WithField("userId", request.UserId).Errorf("Cannot get connections: %v", err)
		return nil, err
	}
	response := &pb.GetResponse{
		Connections: []*pb.Connection{},
	}

	for _, connection := range connections {
		current := mapConnectionToPb(connection)
		response.Connections = append(response.Connections, current)
	}
	return response, nil
}

func (hanler *ConnectionHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	connection := mapPbToConnection(request.Connection)

	newConn, err := hanler.service.Create(connection)
	if err != nil {
		log.Errorf("Can't create connection: %v", err)
		return nil, err
	}

	//prikazi postove za dodatog
	log.Info("Connection created")
	response := &pb.CreateResponse{
		Connection: mapConnectionToPb(newConn),
	}

	return response, nil
}

func (handler *ConnectionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		log.Errorf("Cannot delete connection: %v", err)
		return nil, err
	}
	// obrisati postove kad se obrise konekcija
	log.Info("Connection deleted")
	return &pb.DeleteResponse{}, nil
}

func (handler *ConnectionHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	conn, err := handler.service.Update(request.Id)
	if err != nil {
		log.WithField("connectionId", request.Id).Errorf("Cannot update connection: %v", err)
		return nil, err
	}
	// update postove
	log.WithField("connectionId", request.Id).Infof("Connection updated")
	response := &pb.UpdateResponse{
		Connection: mapConnectionToPb(conn),
	}

	return response, nil
}

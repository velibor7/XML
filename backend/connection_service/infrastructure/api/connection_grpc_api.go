package api

import (
	"context"

	pb "github.com/velibor7/XML/common/proto/connection_service"
	"github.com/velibor7/XML/connection_service/application"
)

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
		return nil, err
	}

	//prikazi postove za dodatog

	response := &pb.CreateResponse{
		Connection: mapConnectionToPb(newConn),
	}

	return response, nil
}

func (handler *ConnectionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	// obrisati postove kad se obrise konekcija

	return &pb.DeleteResponse{}, nil
}

func (handler *ConnectionHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	conn, err := handler.service.Update(request.Id)
	if err != nil {
		return nil, err
	}
	// update postove

	response := &pb.UpdateResponse{
		Connection: mapConnectionToPb(conn),
	}

	return response, nil
}

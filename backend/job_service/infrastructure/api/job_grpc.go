package api

import (
	"context"

	pb "github.com/velibor7/XML/common/proto/job_service"
	"github.com/velibor7/XML/job_service/application"
)

type JobHandler struct {
	pb.UnimplementedJobServiceServer
	service *application.JobService
}

func NewJobHandler(service *application.JobService) *JobHandler {
	return &JobHandler{
		service: service,
	}
}

func (handler *JobHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	Job, err := handler.service.Get(request.Id)
	if err != nil {
		return nil, err
	}
	JobPb := mapJob(Job)

	return &pb.GetResponse{
		Job: JobPb,
	}, nil
}

func (handler *JobHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Jobs, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Job: []*pb.Job{},
	}

	for _, Job := range Jobs {
		temp := mapJob(Job)
		response.Job = append(response.Job, temp)
	}
	return response, nil
}

func (hanlder *JobHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	job := mapPb(request.Job)
	err := hanlder.service.Create(job)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Job: mapJob(job),
	}, nil
}

// func (handler *JobHandler) GetByTitle(ctx context.Context, request *pb.GetByTitleRequest) (*pb.GetByTitleResponse, error) {
// 	Jobs, err := handler.service.GetByTitle(request.Title)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := &pb.GetByTitleResponse{
// 		Job: []*pb.Job{},
// 	}

// 	for _, Job := range Jobs {
// 		temp := mapJob(Job)
// 		response.Job = append(response.Job, temp)
// 	}
// 	return response, nil
// }

package api

import (
	"context"

	"github.com/velibor7/XML/common/loggers"
	pb "github.com/velibor7/XML/common/proto/job_service"
	"github.com/velibor7/XML/common/tracer"
	"github.com/velibor7/XML/job_service/application"
)

var log = loggers.NewJobLogger()

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
	span := tracer.StartSpanFromContext(ctx, "Get")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	Job, err := handler.service.Get(request.Id)
	if err != nil {
		log.WithField("jobId", request.Id).Errorf("Cannot get profile: %v", err)
		return nil, err
	}
	JobPb := mapJob(Job)

	return &pb.GetResponse{
		Job: JobPb,
	}, nil
}

func (handler *JobHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	Jobs, err := handler.service.GetAll()
	if err != nil {
		log.Errorf("Cannot get all jobs: %v", err)
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
	span := tracer.StartSpanFromContext(ctx, "Create")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	job := mapPb(request.Job)
	err := hanlder.service.Create(job)
	if err != nil {
		log.Errorf("Cannot create job: %v", err)
		return nil, err
	}
	log.Info("Job Created")
	return &pb.CreateResponse{
		Job: mapJob(job),
	}, nil
}

func (handler *JobHandler) GetRecommendedJobs(ctx context.Context, request *pb.GetRecommendedJobsRequest) (*pb.GetRecommendedJobsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetRecommended")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	Jobs, err := handler.service.GetRecommendedJobs(request.Id)
	if err != nil {
		log.Errorf("Cannot get recommended jobs: %v", err)
		return nil, err
	}
	response := &pb.GetRecommendedJobsResponse{
		Job: []*pb.Job{},
	}

	for _, Job := range Jobs {
		temp := mapJob(Job)
		response.Job = append(response.Job, temp)
	}
	return response, nil
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

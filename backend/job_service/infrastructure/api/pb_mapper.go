package api

import (
	pb "github.com/velibor7/XML/common/proto/job_service"
	"github.com/velibor7/XML/job_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapJob(job *domain.Job) *pb.Job {
	jobPb := &pb.Job{
		Id:          job.Id.Hex(),
		Title:       job.Title,
		Description: job.Description,
		UserId:      job.UserId.Hex(),
		Skills:      make([]string, 0),
	}
	for _, skill := range job.Skills {
		jobPb.Skills = append(jobPb.Skills, skill)
	}
	return jobPb
}

func mapPb(pbJob *pb.Job) *domain.Job {
	job := &domain.Job{
		Id:          getObjectId(pbJob.Id),
		Title:       pbJob.Title,
		Description: pbJob.Description,
		UserId:      getObjectId(pbJob.UserId),
		Skills:      make([]string, 0),
	}
	for _, skill := range pbJob.Skills {
		job.Skills = append(job.Skills, skill)
	}
	return job
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

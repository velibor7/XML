package api

import (
	pb "github.com/velibor7/XML/common/proto/job_service"
	"github.com/velibor7/XML/job_service/domain"
)

func mapJob(job *domain.Job) *pb.Job {
	jobPb := &pb.Job{
		Id:          job.Id,
		Title:       job.Title,
		Description: job.Description,
		UserId:      job.UserId,
		Skills:      make([]string, 0),
	}
	for _, skill := range job.Skills {
		jobPb.Skills = append(jobPb.Skills, skill)
	}
	return jobPb
}

func mapPb(pbJob *pb.Job) *domain.Job {
	job := &domain.Job{
		Id:          pbJob.Id,
		Title:       pbJob.Title,
		Description: pbJob.Description,
		UserId:      pbJob.UserId,
		Skills:      make([]string, 0),
	}
	for _, skill := range pbJob.Skills {
		job.Skills = append(job.Skills, skill)
	}
	return job
}

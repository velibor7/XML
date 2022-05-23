package application

import (
	"github.com/velibor7/XML/job_service/domain"
)

type JobService struct {
	jobs domain.JobInterface
}

func NewJobService(jobs domain.JobInterface) *JobService {
	return &JobService{
		jobs: jobs,
	}
}

func (service JobService) Get(id string) (*domain.Job, error) {
	return service.jobs.Get(id)
}

func (service JobService) GetAll(search string) ([]*domain.Job, error) {
	return service.jobs.GetAll(search)
}

func (service JobService) Create(job *domain.Job) error {
	return service.jobs.Create(job)
}

func (service JobService) DeleteAll() error {
	return service.jobs.DeleteAll()
}

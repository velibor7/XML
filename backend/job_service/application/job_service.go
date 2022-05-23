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

func (service *JobService) Get(id string) (*domain.Job, error) {
	return service.jobs.Get(id)
}

func (service *JobService) GetAll() ([]*domain.Job, error) {
	return service.jobs.GetAll()
}

func (service *JobService) GetByTitle(title string) ([]*domain.Job, error) {
	return service.jobs.GetByTitle(title)
}

func (service *JobService) Create(job *domain.Job) error {
	return service.jobs.Create(job)
}

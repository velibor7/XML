package application

import (
	"strconv"

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
	Id, _ := strconv.Atoi(id)
	return service.jobs.Get(Id)
}

func (service *JobService) GetAll() ([]*domain.Job, error) {
	return service.jobs.GetAll()
}

// func (service *JobService) GetByTitle(title string) ([]*domain.Job, error) {
// 	return service.jobs.GetByTitle(title)
// }

func (service *JobService) Create(job *domain.Job) error {
	return service.jobs.Create(job)
}

func (service *JobService) GetRecommendedJobs(id string) ([]*domain.Job, error) {

	job, _ := service.Get(id)
	return service.jobs.GetRecommendedJobs(job)
}

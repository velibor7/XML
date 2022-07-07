package domain

type JobInterface interface {
	Get(id int) (*Job, error)
	GetAll() ([]*Job, error)
	//GetByTitle(title string) ([]*Job, error)
	Create(job *Job) error
	GetRecommendedJobs(job *Job) ([]*Job, error)
	//DeleteAll() error
}

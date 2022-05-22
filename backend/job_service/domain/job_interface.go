package domain

type JobInterface interface {
	Get(id string) (*Job, error)
	GetAll(search string) ([]*Job, error)
	Create(job *Job) error
	DeleteAll() error
}

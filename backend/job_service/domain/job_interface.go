package domain

type JobInterface interface {
	Get(id string) (*Job, error)
	GetAll() ([]*Job, error)
	GetByTitle(title string) ([]*Job, error)
	Create(job *Job) error
	DeleteAll() error
}

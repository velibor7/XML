package domain

type ProfileInterface interface {
	Get(id string) (*Profile, error)
	GetAll() ([]*Profile, error)
	Create(profile *Profile) error
	Update(id string, profile *Profile) error
	DeleteAll() error
	Delete(id string) error
}

package domain

type ProfileInterface interface {
	Get(username string) (*Profile, error)
	GetAll() ([]*Profile, error)
	Create(profile *Profile) error
	Update(username string, profile *Profile) error
	DeleteAll() error
}

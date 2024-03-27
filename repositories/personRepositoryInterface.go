package repositories

import "github.com/lordofthemind/gormGinGo/types"

// PersonRepository defines the interface for database operations related to persons.
type PersonRepositoryInterface interface {
	CreatePerson(person *types.PersonType) error
	GetPerson(id string) (*types.PersonType, error)
	GetAllPersons() ([]types.PersonType, error)
	UpdatePerson(person *types.PersonType) error
	DeletePerson(id string) error
}

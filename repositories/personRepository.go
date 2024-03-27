package repositories

import (
	"github.com/lordofthemind/gormGinGo/types"
	"gorm.io/gorm"
)

// PersonRepository defines the interface for database operations related to persons.
type PersonRepository interface {
	CreatePerson(person *types.PersonType) error
	GetPerson(id string) (*types.PersonType, error)
	GetAllPersons() ([]types.PersonType, error)
	UpdatePerson(person *types.PersonType) error
	DeletePerson(id string) error
}

// PersonRepositoryImpl implements PersonRepository.
type PersonRepositoryImpl struct {
	db *gorm.DB
}

// NewPersonRepository creates a new instance of PersonRepositoryImpl.
func NewPersonRepository(db *gorm.DB) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{db: db}
}

func (r *PersonRepositoryImpl) CreatePerson(person *types.PersonType) error {
	return r.db.Create(person).Error
}

func (r *PersonRepositoryImpl) GetPerson(id string) (*types.PersonType, error) {
	var person types.PersonType
	if err := r.db.First(&person, id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepositoryImpl) GetAllPersons() ([]types.PersonType, error) {
	var persons []types.PersonType
	if err := r.db.Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

func (r *PersonRepositoryImpl) UpdatePerson(person *types.PersonType) error {
	return r.db.Save(person).Error
}

func (r *PersonRepositoryImpl) DeletePerson(id string) error {
	return r.db.Delete(&types.PersonType{}, id).Error
}

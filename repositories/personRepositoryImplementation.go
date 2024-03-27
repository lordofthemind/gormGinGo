package repositories

import (
	"github.com/lordofthemind/gormGinGo/types"
	"gorm.io/gorm"
)

// PersonRepository implements the PersonRepositoryInterface.
type PersonRepository struct {
	db *gorm.DB
}

// NewPersonRepository creates a new instance of PersonRepository.
func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

// CreatePerson creates a new person record in the database.
func (r *PersonRepository) CreatePerson(person *types.PersonType) error {
	return r.db.Create(person).Error
}

// GetPerson retrieves a person record from the database by ID.
func (r *PersonRepository) GetPerson(id string) (*types.PersonType, error) {
	var person types.PersonType
	if err := r.db.First(&person, id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

// GetAllPersons retrieves all person records from the database.
func (r *PersonRepository) GetAllPersons() ([]types.PersonType, error) {
	var persons []types.PersonType
	if err := r.db.Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

// UpdatePerson updates a person record in the database.
func (r *PersonRepository) UpdatePerson(person *types.PersonType) error {
	return r.db.Save(person).Error
}

// DeletePerson deletes a person record from the database by ID.
func (r *PersonRepository) DeletePerson(id string) error {
	return r.db.Delete(&types.PersonType{}, id).Error
}

package repositories

import (
	"github.com/lordofthemind/gormGinGo/types"
	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

// CreatePerson creates a new person record in the database.
func (r *PersonRepository) CreatePerson(person *types.PersonType) (*types.PersonType, error) {
	if err := r.db.Create(person).Error; err != nil {
		return nil, err
	}
	return person, nil
}

// GetPerson retrieves a person record from the database based on the given ID.
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

// UpdatePerson updates the details of an existing person record in the database.
func (r *PersonRepository) UpdatePerson(person *types.PersonType) error {
	return r.db.Save(person).Error
}

// DeletePerson deletes a person record from the database based on the given ID.
func (r *PersonRepository) DeletePerson(id string) error {
	return r.db.Delete(&types.PersonType{}, id).Error
}

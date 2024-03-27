package types

import (
	"gorm.io/gorm"
)

type BaseType struct {
	gorm.Model
	IsActive  bool `gorm:"not null;default:false" json:"is_active"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
	DeletedBy uint `json:"deleted_by"`
}

type PersonType struct {
	BaseType
	Username string `gorm:"not null;unique" json:"username" binding:"required"`
	Email    string `gorm:"not null;unique" json:"email" validate:"required,email"`
	Phone    string `gorm:"not null;unique" json:"phone" binding:"required,min=10,max=13"`
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Address  string `gorm:"not null" json:"address"`
	Age      uint   `json:"age"`
	Gender   string `json:"gender"`
	// Add more fields as needed
}

// CreatePerson creates a new person record in the database
func (p *PersonType) CreatePerson(db *gorm.DB) error {
	return db.Create(p).Error
}

// GetPerson retrieves a person record from the database
func (p *PersonType) GetPerson(db *gorm.DB, id string) error {
	return db.First(p, id).Error
}

// GetAllPersons retrieves all person records from the database
func GetAllPersons(db *gorm.DB) ([]PersonType, error) {
	var persons []PersonType
	err := db.Find(&persons).Error
	return persons, err
}

// UpdatePerson updates a person record in the database
func (p *PersonType) UpdatePerson(db *gorm.DB) error {
	return db.Save(p).Error
}

// DeletePerson deletes a person record from the database
func (p *PersonType) DeletePerson(db *gorm.DB) error {
	return db.Delete(p).Error
}

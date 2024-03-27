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

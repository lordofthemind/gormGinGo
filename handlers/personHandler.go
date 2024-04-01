package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGinGo/repositories"
	"github.com/lordofthemind/gormGinGo/types"
	"gorm.io/gorm"
)

// PersonHandlerInterface defines the interface for person handlers.
type PersonHandlerInterface interface {
	CreatePersonHandler(c *gin.Context)
	GetPersonHandler(c *gin.Context)
	GetAllPersonsHandler(c *gin.Context)
	UpdatePersonHandler(c *gin.Context)
	DeletePersonHandler(c *gin.Context)
}

type PersonHandler struct {
	db         *gorm.DB
	personRepo *repositories.PersonRepository
}

func NewPersonHandler(db *gorm.DB) PersonHandlerInterface {
	personRepo := repositories.NewPersonRepository(db)
	return &PersonHandler{db: db, personRepo: personRepo}
}

func (ph *PersonHandler) CreatePersonHandler(c *gin.Context) {
	var person types.PersonType
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid JSON payload"})
		return
	}

	// Call CreatePerson method on personRepo
	createdPerson, err := ph.personRepo.CreatePerson(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to create person"})
		return
	}

	c.JSON(http.StatusCreated, createdPerson)
}

func (ph *PersonHandler) GetPersonHandler(c *gin.Context) {
	id := c.Param("id")
	person, err := ph.personRepo.GetPerson(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (ph *PersonHandler) GetAllPersonsHandler(c *gin.Context) {
	persons, err := ph.personRepo.GetAllPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to retrieve persons"})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func (ph *PersonHandler) UpdatePersonHandler(c *gin.Context) {
	var person types.PersonType
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid JSON payload"})
		return
	}

	if err := ph.personRepo.UpdatePerson(&person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to update person"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (ph *PersonHandler) DeletePersonHandler(c *gin.Context) {
	id := c.Param("id")
	if err := ph.personRepo.DeletePerson(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to delete person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGinGo/repositories"
	"github.com/lordofthemind/gormGinGo/types"
)

// PersonControllerInterface defines the interface for person controllers.
type PersonHandlerInterface interface {
	CreatePersonHandler(c *gin.Context)
	GetPersonHandler(c *gin.Context)
	GetAllPersonsHandler(c *gin.Context)
	UpdatePersonHandler(c *gin.Context)
	DeletePersonHandler(c *gin.Context)
}

type PersonHandler struct {
	personRepo repositories.PersonRepository
}

func NewPersonHandler(personRepo repositories.PersonRepository) *PersonHandler {
	return &PersonHandler{personRepo: personRepo}
}

func (pc *PersonHandler) CreatePersonHandler(c *gin.Context) {
	// Implementation omitted for brevity
	var person types.PersonType
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid JSON payload"})
		return
	}

	if err := pc.personRepo.CreatePerson(&person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to create person"})
		return
	}

	c.JSON(http.StatusCreated, person)
}

func (pc *PersonHandler) GetPersonHandler(c *gin.Context) {
	// Implementation omitted for brevity
	id := c.Param("id")
	person, err := pc.personRepo.GetPerson(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (pc *PersonHandler) GetAllPersonsHandler(c *gin.Context) {
	// Implementation omitted for brevity
	persons, err := pc.personRepo.GetAllPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to retrieve persons"})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func (pc *PersonHandler) UpdatePersonHandler(c *gin.Context) {
	// Implementation omitted for brevity
	var person types.PersonType
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid JSON payload"})
		return
	}

	if err := pc.personRepo.UpdatePerson(&person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to update person"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (pc *PersonHandler) DeletePersonHandler(c *gin.Context) {
	// Implementation omitted for brevity
	id := c.Param("id")
	if err := pc.personRepo.DeletePerson(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to delete person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGinGo/initializers"
	"github.com/lordofthemind/gormGinGo/types"
	"net/http"
)

// CreatePersonHandler creates a new person.
func CreatePersonHandler(ctx *gin.Context) {
	var person types.PersonType

	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid JSON payload"})
		return
	}

	if err := person.CreatePerson(initializers.DB); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to create person"})
		return
	}

	ctx.JSON(http.StatusCreated, person)
}

// GetPersonHandler retrieves a person by ID.
func GetPersonHandler(ctx *gin.Context) {
	var person types.PersonType

	if err := person.GetPerson(initializers.DB, ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "There is no person by this ID."})
		return
	}

	ctx.JSON(http.StatusOK, person)
}

// GetAllPersonsHandler retrieves all persons.
func GetAllPersonsHandler(ctx *gin.Context) {
	persons, err := types.GetAllPersons(initializers.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to find any person"})
		return
	}

	ctx.JSON(http.StatusOK, persons)
}

// UpdatePersonHandler updates a person by ID.
func UpdatePersonHandler(ctx *gin.Context) {
	var person types.PersonType

	if err := person.GetPerson(initializers.DB, ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid JSON payload"})
		return
	}

	if err := person.UpdatePerson(initializers.DB); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to update person"})
		return
	}

	ctx.JSON(http.StatusOK, person)
}

// DeletePersonHandler deletes a person by ID.
func DeletePersonHandler(ctx *gin.Context) {
	var person types.PersonType

	if err := person.GetPerson(initializers.DB, ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "message": "Person not found"})
		return
	}

	if err := person.DeletePerson(initializers.DB); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to delete person"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": person, "message": "Person deleted"})
}

package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGinGo/handlers"
	"github.com/lordofthemind/gormGinGo/initializers"
)

func Run() {
	// Initialize Gin router
	router := gin.Default()

	// Use middleware if needed (e.g., logging, authentication)
	// router.Use(gin.Logger())

	// Initialize repositories
	personRepo := initializers.CreatePersonRepository()

	// Initialize handlers with repositories
	personHandler := handlers.NewPersonHandler(personRepo)

	// Group routes if necessary
	v1 := router.Group("/api/v1")
	{
		v1.POST("/persons", personHandler.CreatePersonHandler)
		v1.GET("/persons/:id", personHandler.GetPersonHandler)
		v1.GET("/persons", personHandler.GetAllPersonsHandler)
		v1.PUT("/persons/:id", personHandler.UpdatePersonHandler)
		v1.DELETE("/persons/:id", personHandler.DeletePersonHandler)
	}

	// Define a port with a fallback value
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090" // Default to port 9090 if PORT environment variable is not set
	}

	// Start the server
	address := fmt.Sprintf(":%s", port)
	if err := router.Run(address); err != nil {
		// Handle error if router fails to start
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}

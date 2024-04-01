package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGinGo/handlers"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	// Initialize Gin router
	router := gin.Default()

	personHandler := handlers.NewPersonHandler(db)

	// Routes
	v1 := router.Group("/api/v1")
	{
		// Person routes
		persons := v1.Group("/persons")
		{
			persons.POST("/", personHandler.CreatePersonHandler)
			persons.GET("/:id", personHandler.GetPersonHandler)
			persons.GET("/", personHandler.GetAllPersonsHandler)
			persons.PUT("/:id", personHandler.UpdatePersonHandler)
			persons.DELETE("/:id", personHandler.DeletePersonHandler)
		}
	}

	// Start server
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

package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGinGo/handlers"
)

func Run() {
	// Initialize Gin router
	router := gin.Default()

	// Use middleware if needed (e.g., logging, authentication)
	// router.Use(gin.Logger())

	// Group routes if necessary
	v1 := router.Group("/api/v1")
	{
		v1.POST("/persons", handlers.CreatePersonHandler)
		v1.GET("/persons/:id", handlers.GetPersonHandler)
		v1.GET("/persons", handlers.GetAllPersonsHandler)
		v1.PUT("/persons/:id", handlers.UpdatePersonHandler)
		v1.DELETE("/persons/:id", handlers.DeletePersonHandler)
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

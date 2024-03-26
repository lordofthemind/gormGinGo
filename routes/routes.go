package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGinGo/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgresql()
	initializers.SyncPostgresql()
}

func Run() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.Run("localhost:" + os.Getenv("PORT"))
}

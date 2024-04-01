package main

import (
	"github.com/lordofthemind/gormGinGo/initializers"
	"github.com/lordofthemind/gormGinGo/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgresql()
	// initializers.CreatePersonRepository()
	initializers.SyncPostgresql()
}

func main() {
	routes.Run(initializers.DB)
}

package main

import (
	"github.com/lordofthemind/gormGinGo/initializers"
	"github.com/lordofthemind/gormGinGo/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgresql()
	initializers.SyncPostgresql()
}

func main() {
	routes.Run()
}

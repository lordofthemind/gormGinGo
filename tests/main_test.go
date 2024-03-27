package tests

import (
	"github.com/lordofthemind/gormGinGo/initializers"
)

func init() {
	initializers.ConnectToPostgresql()
	initializers.LoadEnvVariables()

}

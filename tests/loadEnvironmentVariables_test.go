package tests

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariablesTest() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	fmt.Println("Environment variables loaded")
	return nil
}

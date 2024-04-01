package tests

import (
	"fmt"
	"log"

	"github.com/lordofthemind/gormGinGo/types"
)

// SyncPostgresql synchronizes the PersonType model with the PostgreSQL database.
func SyncPostgresqlTest() error {
	err := MOCK_DB.AutoMigrate(&types.PersonType{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate PersonType model in test environment: %w", err)
	}

	log.Println("Synchronized the PersonType model")
	log.Println("Synchronized the test database")

	return nil
}

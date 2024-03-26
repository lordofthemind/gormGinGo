package initializers

import (
	"fmt"
	"log"

	"github.com/lordofthemind/gormGinGo/types"
)

// SyncPostgresql synchronizes the PersonType model with the PostgreSQL database.
func SyncPostgresql() error {
	err := DB.AutoMigrate(&types.PersonType{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate PersonType model: %w", err)
	}

	log.Println("Synchronized the PersonType model")
	log.Println("Synchronized the database")

	return nil
}

package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/lordofthemind/gormGinGo/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgresql() {
	var err error
	postgresqlConnectionString := os.Getenv("PG_DB_CONNECTION")
	DB, err = gorm.Open(postgres.Open(postgresqlConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
}

func CreatePersonRepository() repositories.PersonRepository {
	return repositories.NewPersonRepository(DB)
}

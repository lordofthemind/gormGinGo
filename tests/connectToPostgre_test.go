package tests

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var MOCK_DB *gorm.DB

func ConnectToPostgresqlTest() {
	var err error
	// postgresqlConnectionString := os.Getenv("PG_DB_CONNECTION_TEST")
	postgresqlConnectionString := "postgres://postgres:secret-gorm@localhost:5432/person_type_gorm_test?sslmode=disable"

	MOCK_DB, err = gorm.Open(postgres.Open(postgresqlConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to test database")
}

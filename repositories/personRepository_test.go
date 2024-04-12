package repositories

import (
	"fmt"
	"testing"

	"github.com/lordofthemind/gormGinGo/types"
)

func TestCreatePerson(t *testing.T) {
	tests.ConnectToPostgresqlTest()
	defer tests.MOCK_DB.Migrator().DropTable(&types.PersonType{})
	_ = NewPersonRepository(tests.MOCK_DB)
	fmt.Print("db connected")

}

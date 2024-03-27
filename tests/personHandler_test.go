package tests

import (
	"context"
	"testing"

	"github.com/lordofthemind/gormGinGo/handlers"
	"github.com/lordofthemind/gormGinGo/helpers"
	"github.com/stretchr/testify/require"
)

func TestCreatePersonHandler(t *testing.T) {
	t.Log("Testing CreatePersonHandler")

	rg := helpers.NewRandomGenerator()
	arg := CreatePersonParams{
		Username: rg.RandomUsername(),
		Email:    rg.RandomEmail(),
		Phone:    rg.RandomPhone(),
		Name:     rg.RandomName(),
		Address:  rg.RandomAddress(),
		Age:      rg.RandomAge(),
		Gender:   rg.RandomGender(),
	}

	// Create a new person
	person, err := handlers.CreatePersonHandler(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, arg.Owner, person.Owner)
	require.Equal(t, arg.Balance, person.Balance)
	require.Equal(t, arg.Currency, person.Currency)
	require.NotZero(t, person.ID)
	require.NotZero(t, person.CreatedAt)

	return person
}

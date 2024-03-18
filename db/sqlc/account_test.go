package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mod/utils"
)

func createRandomAccount(t *testing.T) Account {
	
	arg := CreateAccountParams{
		
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testqueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}




func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}



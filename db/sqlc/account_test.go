package db

import (
	"context"
	"testing"
	"time"

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
func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testqueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}



func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}



package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	fmt.Println("An Starting TestCreateAccount")
	arg := CreateAccountParams{
		Username: "r9",
		Password: "123ddd",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	arg.Username = "r99"

	account, err = testQueries.CreateAccount(context.Background(), arg)
	account2, err2 := testQueries.GetAccount(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NoError(t, err2)
	require.NotEmpty(t, account2)
}

func TestCreate2Accounts(t *testing.T) {
	store := NewStore(testDb)
	store.Create2AccountsTx(context.Background())
}

func TestCreateAccount2(t *testing.T) {
	fmt.Println("An Starting TestCreateAccount2")
	arg := CreateAccountParams{
		Username: "fgggg",
		Password: "123ddd",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
}

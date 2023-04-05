package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    faker.FirstName(),
		Balance:  faker.UnixTime(),
		Currency: faker.Currency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
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

func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	fetchedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)
	require.Equal(t, createdAccount.ID, fetchedAccount.ID)
	require.Equal(t, createdAccount.Owner, fetchedAccount.Owner)
	require.Equal(t, createdAccount.Balance, fetchedAccount.Balance)
	require.Equal(t, createdAccount.Currency, fetchedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, fetchedAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: faker.UnixTime(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, createdAccount.ID, updatedAccount.ID)
	require.Equal(t, createdAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, createdAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)

	fetchedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fetchedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

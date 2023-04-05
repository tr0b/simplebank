package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        faker.UnixTime(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer

}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)
	fetchedTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedTransfer)
	require.Equal(t, createdTransfer.ID, fetchedTransfer.ID)
	require.Equal(t, createdTransfer.Amount, fetchedTransfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, fetchedTransfer.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)
	arg := UpdateTransferParams{
		ID:     createdTransfer.ID,
		Amount: faker.UnixTime(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)

	require.Equal(t, createdTransfer.ID, updatedTransfer.ID)
	require.Equal(t, updatedTransfer.Amount, arg.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, updatedTransfer.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)

	fetchedTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fetchedTransfer)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

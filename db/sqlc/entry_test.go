package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    faker.UnixTime(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry

}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	createdEntry := createRandomEntry(t)
	fetchedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedEntry)
	require.Equal(t, createdEntry.ID, fetchedEntry.ID)
	require.Equal(t, createdEntry.Amount, fetchedEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, fetchedEntry.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	createdEntry := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     createdEntry.ID,
		Amount: faker.UnixTime(),
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)

	require.Equal(t, createdEntry.ID, updatedEntry.ID)
	require.Equal(t, updatedEntry.Amount, arg.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, updatedEntry.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	createdEntry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)

	fetchedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fetchedEntry)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

package db

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"

	"github.com/tr0b/simplebank/util"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(faker.Password())
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       faker.Username(),
		HashedPassword: hashedPassword,
		FullName:       faker.Name(),
		Email:          faker.Email(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)
	fetchedUser, err := testQueries.GetUser(context.Background(), createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedUser)
	require.Equal(t, createdUser.Username, fetchedUser.Username)
	require.Equal(t, createdUser.HashedPassword, fetchedUser.HashedPassword)
	require.Equal(t, createdUser.FullName, fetchedUser.FullName)
	require.Equal(t, createdUser.Email, fetchedUser.Email)
	require.WithinDuration(t, createdUser.PasswordChangedAt, fetchedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, createdUser.CreatedAt, fetchedUser.CreatedAt, time.Second)
}

// func TestUpdateUser(t *testing.T) {
// 	createdUser := createRandomUser(t)
// 	arg := UpdateUserParams{
// 		ID:      createdUser.ID,
// 		Balance: faker.UnixTime(),
// 	}
//
// 	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, updatedUser)
//
// 	require.Equal(t, createdUser.ID, updatedUser.ID)
// 	require.Equal(t, createdUser.Owner, updatedUser.Owner)
// 	require.Equal(t, arg.Balance, updatedUser.Balance)
// 	require.Equal(t, createdUser.Currency, updatedUser.Currency)
// 	require.WithinDuration(t, createdUser.CreatedAt, updatedUser.CreatedAt, time.Second)
// }
//
// func TestDeleteUser(t *testing.T) {
// 	createdUser := createRandomUser(t)
// 	err := testQueries.DeleteUser(context.Background(), createdUser.ID)
// 	require.NoError(t, err)
//
// 	fetchedUser, err := testQueries.GetUser(context.Background(), createdUser.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, fetchedUser)
// }
//
// func TestListUsers(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomUser(t)
// 	}
//
// 	arg := ListUsersParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}
//
// 	users, err := testQueries.ListUsers(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, users, 5)
//
// 	for _, user := range users {
// 		require.NotEmpty(t, user)
// 	}
// }

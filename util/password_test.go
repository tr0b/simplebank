package util

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := faker.Password()

	// check hashed password matches
	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	// checks wrong password does not match with hash
	wrongPassword := faker.Password()
	err = CheckPassword(wrongPassword, hashedPassword1)

	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// check re-hashed password is not the same
	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}

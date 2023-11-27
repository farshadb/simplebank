package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(10)

	hashedPassword1, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = VerifyPassword(password, hashedPassword1)
	require.NoError(t, err)

	//
	wrongPassword := RandomString(10)
	err = VerifyPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)

}

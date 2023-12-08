package test

import (
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := util.RandomString(8)

	hash, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = util.CheckPassword(password, hash)
	require.NoError(t, err)

	wrongPassword := util.RandomString(6)
	err = util.CheckPassword(wrongPassword, hash)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
package test

import (
	"context"
	db "simple-bank/db/sqlc"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func creatRandomUser(t *testing.T) db.User {
	hashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := db.CreatUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testingQueries.CreatUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	creatRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := creatRandomUser(t)
	user2, err := testingQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

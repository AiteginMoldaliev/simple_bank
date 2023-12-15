package test

import (
	"simple-bank/token"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := token.NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issudeAt := time.Now()
	expiredAt := issudeAt.Add(duration)

	pasetoToken, payload, err := maker.CraeteToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, pasetoToken)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifiToken(pasetoToken)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issudeAt, payload.IssudeAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := token.NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	pasetoToken, payload, err := maker.CraeteToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, pasetoToken)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifiToken(pasetoToken)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrExpiredToken.Error())
	require.Nil(t, payload)
}
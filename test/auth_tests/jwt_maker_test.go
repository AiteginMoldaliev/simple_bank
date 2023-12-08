package test

import (
	"fmt"
	"simple-bank/token"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := token.NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issudeAt := time.Now()
	expiredAt := issudeAt.Add(duration)

	jwtToken, err := maker.CraeteToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, jwtToken)

	claims, err := maker.VerifiToken(jwtToken)
	require.NoError(t, err)
	require.NotEmpty(t, claims)

	require.NotZero(t, claims)
	require.Equal(t, username, claims["username"])

	claimIssuedAt := claims["issudeAt"]
	parsedIssuedAt, err := time.Parse(time.RFC3339, fmt.Sprint(claimIssuedAt))
	require.NoError(t, err)
	claimExpiredAt := claims["expiredAt"]
	parsedExpiredAt, err := time.Parse(time.RFC3339, fmt.Sprint(claimExpiredAt))
	require.NoError(t, err)
	require.WithinDuration(t, issudeAt, parsedIssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, parsedExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := token.NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	jwtToken, err := maker.CraeteToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, jwtToken)

	claims, err := maker.VerifiToken(jwtToken)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrExpiredToken.Error())
	require.Nil(t, claims)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := token.NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"id":        payload.ID,
		"username":  payload.Username,
		"issudeAt":  payload.IssudeAt,
		"expiredAt": payload.ExpiredAt,
	})

	tk, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := token.NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	claims, err := maker.VerifiToken(tk)
	require.Error(t, err)
	require.EqualError(t, err, "token is unverifiable: error while executing keyfunc: token is invalid")
	require.Nil(t, claims)
}
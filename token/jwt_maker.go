package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (JwtMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

func (maker *JWTMaker) CraeteToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        payload.ID,
		"username":  payload.Username,
		"issudeAt":  payload.IssudeAt,
		"expiredAt": payload.ExpiredAt,
	})
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) VerifiToken(token string) (jwt.MapClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		if _, ok := claims["id"].(string); !ok {
			return nil, ErrInvalidToken
		}
		if _, ok := claims["username"].(string); !ok {
			return nil, ErrInvalidToken
		}
		if _, ok := claims["issudeAt"].(string); !ok {
			return nil, ErrInvalidToken
		}
		expiredAt, ok := claims["expiredAt"].(string) 
		if !ok {
			return nil, ErrInvalidToken
		}
		parsedExpiredAt, err := time.Parse(time.RFC3339, expiredAt)
		if err != nil {
			return nil, fmt.Errorf("wrong time format")
		}
		if time.Now().After(parsedExpiredAt) {
			return nil, ErrExpiredToken
		}

		return claims, nil
	}

	return nil, ErrInvalidToken
}

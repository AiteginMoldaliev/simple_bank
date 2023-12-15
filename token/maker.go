package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtMaker interface {
	CraeteToken(username string, duration time.Duration) (string, *Payload, error)
	VerifiToken(token string) (jwt.MapClaims, error)
}

type PMaker interface {
	CraeteToken(username string, duration time.Duration) (string, *Payload, error)
	VerifiToken(token string) (*Payload, error)
} 

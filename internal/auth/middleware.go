package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"simple-bank/token"
	"simple-bank/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func PasetoAuthMiddleware(pasetoMaker token.PMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(util.AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != util.AuthorizationTypeBearer {
			err := fmt.Errorf("invalid authorization type: %v", util.AuthorizationTypeBearer)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		accessToken := fields[1]
		payload, err := pasetoMaker.VerifiToken(accessToken)
		if err != nil {
			err := errors.New("invalid authorization token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set(util.AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}

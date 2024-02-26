package api

import (
	"NDE_backend/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func groupAuthMiddleware(tokenMaker token.Maker, requiredGroup string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("authorization header is not provided")))
			return
		}

		// Extract the token from the Authorization header.
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("invalid authorization header format")))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("access token is invalid: %v", err)))
			return
		}

		// Now, assuming your payload includes something like user's groups or roles,
		// you should check if the user belongs to the requiredGroup.
		if !strings.Contains(payload.UserGroup, requiredGroup) {
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(fmt.Errorf("user does not have the required group rights")))
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/auth"
	"github.com/mag30/project-backend/domain/base"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	//UserIDKey value type is uuid.UUID
	UserIDKey = "userID"
)

// SetAuthorizationCheck adds authorization check to middleware chain.
func SetAuthorizationCheck(JWTManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, base.ResponseFailure{
				Status:     http.StatusText(http.StatusUnauthorized),
				Blame:      base.BlameUser,
				TrackingID: GetTrackingId(c),
				Message:    "unauthorized",
			})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, base.ResponseFailure{
				Status:     http.StatusText(http.StatusUnauthorized),
				Blame:      base.BlameUser,
				TrackingID: GetTrackingId(c),
				Message:    "unauthorized",
			})
			return
		}

		stringUserID, err := JWTManager.Parse(headerParts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, base.ResponseFailure{
				Status:     http.StatusText(http.StatusUnauthorized),
				Blame:      base.BlameUser,
				TrackingID: GetTrackingId(c),
				Message:    "unauthorized",
			})
			return
		}

		userID, err := uuid.Parse(stringUserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, base.ResponseFailure{
				Status:     http.StatusText(http.StatusUnauthorized),
				Blame:      base.BlameUser,
				TrackingID: GetTrackingId(c),
				Message:    "unauthorized",
			})
			return
		}

		c.Set(UserIDKey, userID)
		c.Next()
	}
}

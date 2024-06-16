package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const trackingIdHeader = "X-Tracking-Id"
const sessionIdHeader = "X-Session-Id"

// GetTrackingId returns tracking id from context.
func GetTrackingId(c *gin.Context) string {
	return c.GetString(trackingIdHeader)
}

// GetTrackingIdHeader returns tracking id header name.
func GetTrackingIdHeader() string {
	return trackingIdHeader
}

// GetSessionId returns session id from context.
func GetSessionId(c *gin.Context) string {
	return c.GetString(sessionIdHeader)
}

// GetSessionIdHeader returns sessionid header name.
func GetSessionIdHeader() string {
	return sessionIdHeader
}

// SetTracingContext adds tracing context to middleware chain.
// Tracing context consist of tracking and session IDs.
func SetTracingContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		tId := c.GetHeader(trackingIdHeader)
		sId := c.GetHeader(sessionIdHeader)

		if tId == "" {
			tId = uuid.New().String()
			c.Header(trackingIdHeader, tId)
		}

		if sId == "" {
			sId = uuid.New().String()
			c.Header(sessionIdHeader, sId)
		}

		c.Set(trackingIdHeader, tId)
		c.Set(sessionIdHeader, sId)

		c.Next()
	}
}

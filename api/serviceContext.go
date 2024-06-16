package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mag30/project-backend/api/middleware"
)

type Key string

const (
	trackingIDKey Key = "trackingID"
	sessionIDKey  Key = "sessionID"
)

func GetTrackingIDKey() Key {
	return trackingIDKey
}

func GetSessionIDKey() Key {
	return sessionIDKey
}

func GetExtendedContext(c *gin.Context) context.Context {
	ctx := context.WithValue(c, trackingIDKey, middleware.GetTrackingId(c))
	ctx = context.WithValue(ctx, sessionIDKey, middleware.GetSessionId(c))
	return ctx
}

func GetExtendedSystemContext(ctx context.Context, tracingID string) context.Context {
	ctx = context.WithValue(ctx, trackingIDKey, tracingID)
	return ctx
}

func GetTrackingIDFromContext(ctx context.Context) string {
	interfaceTrackingID := ctx.Value(trackingIDKey)
	if interfaceTrackingID != nil {
		return interfaceTrackingID.(string)
	}
	return ""
}

func GetSessionIDFromContext(ctx context.Context) string {
	return ctx.Value(sessionIDKey).(string)
}

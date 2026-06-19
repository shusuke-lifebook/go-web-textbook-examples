package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ctxKey int

const requestIDKey ctxKey = 0

// RequestID は リクエストIDをヘッダとコンテキストに付与する
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader("X-Request-ID")
		if rid == "" {
			rid = uuid.New().String()
		}
		c.Header("X-Request-ID", rid)

		ctx := context.WithValue(c.Request.Context(), requestIDKey, rid)
		c.Request = c.Request.WithContext(ctx)
		c.Set("request_id", rid)
		c.Next()
	}
}

// RequestIDFromContext は contextに載せたIDを取り出す
func RequestIDFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}
	return ""
}

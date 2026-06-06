// Package middleware
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Sampleはc.Next()の前後で何が起きるかを示すためのひな型
func Sample() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // 後継ミドルウェアとハンドラを実行する

		elapsed := time.Since(start)
		_ = elapsed // 例示のため使わない
	}
}

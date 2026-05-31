// Package router
package router

import (
	"ch01-gin-intro/internal/handler"

	"github.com/gin-gonic/gin"
)

// New はルーティング設定済みのエンジンを返す
func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.Ping)
	return r
}

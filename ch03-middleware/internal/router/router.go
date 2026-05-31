// Package router
package router

import (
	"ch03-middleware/internal/handler"

	"github.com/gin-gonic/gin"
)

// New はルーティング設定済みのエンジンを返す
func New(taskHandler *handler.TaskHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", taskHandler.Create)
			tasks.GET("", taskHandler.List)
			tasks.GET("/:id", taskHandler.Get)
			tasks.PATCH("/:id", taskHandler.Update)
			tasks.DELETE("/:id", taskHandler.Delete)
		}
	}
	return r
}

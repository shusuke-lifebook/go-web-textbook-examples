// Package router
package router

import (
	"ch04-postgres-sqlc/internal/handler"
	mw "ch04-postgres-sqlc/internal/middleware"
	"log/slog"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Deps struct {
	Logger      *slog.Logger
	RateLimiter *mw.IPRateLimiter
	TaskHandler *handler.TaskHandler
	Production  bool
}

func New(d Deps) *gin.Engine {
	r := gin.New()

	// healthz はミドルウェアチェーンの前に登録する。CORS / レート制限 / Gzip
	// の影響を受けずに常に 200 を返せるようにし、Cloud Run や Kubernetes の
	// プローブが確実に疎通するようにする。
	r.GET("/healthz", func(c *gin.Context) {
		c.String(200, "OK")
	})

	r.Use(mw.Recovery(d.Logger))
	r.Use(mw.RequestID())
	r.Use(mw.Logger(d.Logger))
	r.Use(mw.SecurityHeaders(d.Production))
	r.Use(cors.New(corsConfig()))
	r.Use(d.RateLimiter.Middleware())
	r.Use(gzipMiddleware())

	v1 := r.Group("/api/v1")
	registerTaskRoutes(v1, d.TaskHandler)
	return r
}

func registerTaskRoutes(g *gin.RouterGroup, h *handler.TaskHandler) {
	tasks := g.Group("/tasks")
	tasks.POST("", h.Create)
	tasks.GET("", h.List)
	tasks.GET("/:id", h.Get)
	tasks.PATCH("/:id", h.Update)
	tasks.DELETE("/:id", h.Delete)
}

// New はルーティング設定済みのエンジンを返す
// func New(taskHandler *handler.TaskHandler) *gin.Engine {
// 	r := gin.Default()

// 	v1 := r.Group("/api/v1")
// 	{
// 		tasks := v1.Group("/tasks")
// 		{
// 			tasks.POST("", taskHandler.Create)
// 			tasks.GET("", taskHandler.List)
// 			tasks.GET("/:id", taskHandler.Get)
// 			tasks.PATCH("/:id", taskHandler.Update)
// 			tasks.DELETE("/:id", taskHandler.Delete)
// 		}
// 	}
// 	return r
// }

func corsConfig() cors.Config {
	return cors.Config{
		AllowOrigins: []string{
			"https://app.example.com",
			"http://localhost:5713",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length", "X-Request-ID",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}

func gzipMiddleware() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{
		".png", ".jpg", "jpeg", "webp", ".pdf", ".mp4",
	}),
		gzip.WithExcludedPaths([]string{"/healthz", "/metrics"}),
		gzip.WithMinLength(1024),
	)
}

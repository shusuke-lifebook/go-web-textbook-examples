package main

import (
	"ch04-postgres-sqlc/internal/handler"
	mw "ch04-postgres-sqlc/internal/middleware"
	"ch04-postgres-sqlc/internal/repository"
	"ch04-postgres-sqlc/internal/router"
	"ch04-postgres-sqlc/internal/usecase"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	repo := repository.NewInMemoryTaskRepo()
	taskUsecase := usecase.New(repo)
	th := handler.NewTaskHandler(taskUsecase)

	limiter := mw.NewIPRateLimiter(rate.Limit(10), 20)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go limiter.StartGC(ctx, 5*time.Minute, 1*time.Hour)

	r := router.New(router.Deps{
		Logger:      logger,
		RateLimiter: limiter,
		TaskHandler: th,
		Production:  os.Getenv("APP_ENV") == "production",
	})
	if err := r.Run(":8080"); err != nil {
		logger.Error("server exited", "error", err)
		os.Exit(1)
	}
}

package main

import (
	"ch02-routing/internal/handler"
	"ch02-routing/internal/repository"
	"ch02-routing/internal/router"
	"ch02-routing/internal/usecase"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	repo := repository.NewInMemoryTaskRepo()
	taskUsecase := usecase.New(repo)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	r := router.New(taskHandler)
	if err := r.Run(":8080"); err != nil {
		logger.Error("server exited", "error", err)
		os.Exit(1)
	}
}

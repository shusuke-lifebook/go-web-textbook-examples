package main

import (
	"ch03-middleware/internal/handler"
	"ch03-middleware/internal/repository"
	"ch03-middleware/internal/router"
	"ch03-middleware/internal/usecase"
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

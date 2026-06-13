package repository

import (
	"ch04-postgres-sqlc/internal/domain"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// mapPgError は pgx / pgconn のエラーを domain エラーに畳み込む
func mapPgError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrTaskNotFound
	}
	return mapPgConnError(err)
}

func mapPgConnError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return domain.ErrDuplicate
		case "23503":
			return domain.ErrForeignKey
		case "23514":
			return domain.ErrCheckViolation
		}
	}
	return fmt.Errorf("db: %w", err)
}

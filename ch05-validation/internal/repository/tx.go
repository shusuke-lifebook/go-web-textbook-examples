package repository

import (
	dbgen "ch05-validation/internal/db/gen"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TxRunner struct {
	pool *pgxpool.Pool
}

func NewTxRunner(pool *pgxpool.Pool) *TxRunner {
	return &TxRunner{pool: pool}
}

// Run は pgx.BeginFunc で Commit / Rollback を自動化する
func (r *TxRunner) Run(ctx context.Context, fn func(ctx context.Context, q *dbgen.Queries) error) error {
	return pgx.BeginFunc(ctx, r.pool, func(tx pgx.Tx) error {
		return fn(ctx, dbgen.New(tx))
	})
}

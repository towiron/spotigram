package interfaces

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type ServiceSyncer interface {
	Sync()
}

type RepositorySpotify interface{}

type DB interface {
	QueryRow(ctx context.Context, query string, args ...any) *sqlx.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

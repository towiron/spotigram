package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/towiron/spotigram/internal/pkg/config"
	"github.com/towiron/spotigram/internal/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"

	"github.com/towiron/spotigram/internal/interfaces"
	"github.com/towiron/spotigram/internal/pkg/global"
)

type Postgres struct {
	Cfg    config.Configer
	DB     *sqlx.DB
	Logger logger.Logger
}

type Options struct {
	fx.In
	fx.Lifecycle
	Config config.Configer
	Logger logger.Logger
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(New, fx.As(new(interfaces.DB))),
	),
)

func New(opts Options) *Postgres {
	var (
		db *sqlx.DB
	)

	postgresDB := &Postgres{
		Cfg:    opts.Config,
		DB:     db,
		Logger: opts.Logger,
	}

	opts.Lifecycle.Append(fx.Hook{
		OnStart: postgresDB.onStart,
		OnStop:  postgresDB.onStop,
	})

	return postgresDB
}

func (p *Postgres) onStart(ctx context.Context) error {
	// Establish the connection to PostgreSQL
	db, err := sqlx.ConnectContext(ctx, "pgx", p.Cfg.String(global.ENV_REPOSITORY_POSTGRES_DSN))
	if err != nil {
		return err
	}

	// Connection configuration
	// more info here https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(p.Cfg.Int(global.ENV_REPOSITORY_POSTGRES_MAX_OPEN_CONNECTIONS))
	db.SetMaxIdleConns(p.Cfg.Int(global.ENV_REPOSITORY_POSTGRES_MAX_IDLE_CONNECTIONS))
	db.SetConnMaxLifetime(p.Cfg.Duration(global.ENV_REPOSITORY_POSTGRES_CONNECTION_MAX_LIFETIME))

	// Assign the connected DB instance to the struct
	p.DB = db
	p.Logger.DebugF("connected to PostgreSQL")

	if err = p.migrateUp(); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) onStop(ctx context.Context) error {
	if p.DB != nil {
		// Gracefully close the database connection
		return p.DB.Close()
	}
	return nil
}

// Ping checks if the PostgreSQL database is available
func (p *Postgres) Ping(ctx context.Context) (err error) {
	p.Logger.Debug("repository.postgres.Ping()")
	if err := p.DB.PingContext(ctx); err != nil {
		p.Logger.DebugF("error while pinging PostgreSQL: %v", err)
		return err
	}

	p.Logger.Debug("PostgreSQL is available")
	return nil
}

func (p *Postgres) migrateUp() error {
	migrationDir := "file://migrations" // Assuming your migrations folder is at root level

	databaseURL := p.Cfg.String(global.ENV_REPOSITORY_POSTGRES_DSN)

	// Initialize a new migration instance
	m, err := migrate.New(migrationDir, databaseURL)
	if err != nil {
		p.Logger.DebugF("Failed to initialize migrations: %v", err)
		return err
	}

	// Run migrations (up)
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		p.Logger.DebugF("Migration failed: %v", err)
		return err
	}

	p.Logger.Debug("Migrations applied successfully!")

	return nil
}

// TODO: migrate down if need

func (p *Postgres) QueryRow(ctx context.Context, query string, args ...any) *sqlx.Row {
	if p.DB == nil {
		p.Logger.Error("database connection is not initialized")
		return nil
	}

	// Logging query execution
	p.Logger.DebugF("Executing QueryRow: %s, args: %v", query, args)

	// Execute the query and return the row
	return p.DB.QueryRowxContext(ctx, query, args...)
}

func (p *Postgres) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if p.DB == nil {
		p.Logger.Error("ExecContext called but database is not initialized")
		return nil, errors.New("database is not initialized")
	}

	p.Logger.DebugF("Executing ExecContext: %s, args: %v", query, args)

	return p.DB.ExecContext(ctx, query, args...)
}

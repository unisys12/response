package response

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/responserms/response/internal/ent"
	_ "github.com/responserms/response/internal/ent/runtime"

	// pgx is used for the postgres driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

func (c *Core) initEnt(ctx context.Context) error {
	if c.config.DatabaseConnURL == nil {
		return errors.New("database configuration is required")
	}

	var client *ent.Client
	var err error

	switch c.config.DatabaseConnURL.Scheme {
	case "postgresql":
		client, err = c.newPostgresClient(ctx)
	case "mysql":
		client, err = c.newMySQLClient(ctx)
	default:
		return errors.New("unsupported driver configured")
	}

	if err != nil {
		return err
	}

	// TODO: Add Debug() logging
	c.ent = client

	return nil
}

func (c *Core) newPostgresClient(ctx context.Context) (*ent.Client, error) {
	db, err := sql.Open("pgx", c.config.DatabaseConnURL.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := c.waitForDatabase(ctx, db, 5); err != nil {
		return nil, fmt.Errorf("failed to Ping database: %w", err)
	}

	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver)), nil
}

func (c *Core) newMySQLClient(ctx context.Context) (*ent.Client, error) {
	return nil, nil
}

func (c *Core) waitForDatabase(ctx context.Context, db *sql.DB, maxRetries int) error {
	var checkTimer = time.NewTicker(1 * time.Second)
	var attempts int

	err := db.Ping()
	if err == nil {
		return nil
	}

	c.logger.Error().
		Err(err).
		Int("maxRetries", maxRetries).
		Msg("failed to Ping database, retrying up to maxRetries attempts at 1 second intervals")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-checkTimer.C:
			err := db.Ping()

			if err != nil && attempts >= maxRetries {
				return err
			}

			c.logger.Error().
				Err(err).
				Int("attempt", attempts+1).
				Msg("failed to Ping database again, retrying in 1 second")

			attempts++
		}
	}
}

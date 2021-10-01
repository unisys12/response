package response

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/responserms/response/internal/ent/migrate"
)

type SchemaChangeConfig struct {
	NoDropColumns bool
}

var ErrSchemaChangesAvailable = errors.New("schema changes required")

func (c *Core) getMigrateOptions(config *SchemaChangeConfig) []schema.MigrateOption {
	if config == nil {
		config = &SchemaChangeConfig{}
	}

	var migrateOptions = []schema.MigrateOption{
		migrate.WithDropIndex(true),
		migrate.WithForeignKeys(true),
		migrate.WithGlobalUniqueID(true),
	}

	if config.NoDropColumns {
		migrateOptions = append(migrateOptions, migrate.WithDropColumn(true))
	}

	return migrateOptions
}

func (c *Core) HasSchemaChanges(ctx context.Context) (bool, error) {
	const (
		// emptyTxString provides the standard output for an empty SQL transaction for which
		// no statements were generated. This means that no changes are needed to the schema.
		//
		// Hopefully Ent exposes this in the future as a first-class option. But for now, we
		// hack it.
		emptyTxString = `BEGIN;
COMMIT;`
	)

	var buf bytes.Buffer
	if err := c.ent.Schema.WriteTo(ctx, &buf, c.getMigrateOptions(&SchemaChangeConfig{})...); err != nil {
		return false, fmt.Errorf("unable to determine schema changes: %w", err)
	}

	switch c.config.DatabaseConnURL.Scheme {
	case "postgresql":
		fallthrough
	case "mysql":
		if strings.Trim(buf.String(), "\n") == emptyTxString {
			return false, nil
		}
	}

	return true, nil
}

func (c *Core) ApplySchemaChanges(ctx context.Context, cfg *SchemaChangeConfig) error {
	err := c.ent.Schema.Create(ctx, c.getMigrateOptions(cfg)...)

	if err != nil {
		return fmt.Errorf("db.Schema.Create: %w", err)
	}

	return nil
}

func (c *Core) WriteSchemaChanges(ctx context.Context, cfg *SchemaChangeConfig, out io.Writer) error {
	err := c.ent.Schema.WriteTo(ctx, out, c.getMigrateOptions(cfg)...)

	if err != nil {
		return fmt.Errorf("db.Schema.WriteTo: %w", err)
	}

	return nil
}

package database

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"go.uber.org/fx"

	"github.com/Willsem/health-metrics/internal/generated/ent"
)

func NewClient(lc fx.Lifecycle, c *Config) (*ent.Client, error) {
	driver, err := sql.Open(dialect.Postgres, c.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("open sql connection driver: %w", err)
	}

	err = driver.DB().Ping()
	if err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	client := ent.NewClient(ent.Driver(driver))
	lc.Append(
		fx.Hook{
			OnStop: func(_ context.Context) error {
				return client.Close()
			},
		},
	)

	return client, nil
}

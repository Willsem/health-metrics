package database

import (
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"

	"github.com/Willsem/health-metrics/internal/generated/ent"
)

func NewClient(c *Config) (*ent.Client, error) {
	driver, err := sql.Open(dialect.Postgres, c.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("open sql connection driver: %w", err)
	}

	err = driver.DB().Ping()
	if err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return ent.NewClient(ent.Driver(driver)), nil
}

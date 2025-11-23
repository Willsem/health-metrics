package postgres

import (
	"context"
	"fmt"
	"health-metrics/internal/generated/ent"

	_ "github.com/lib/pq"
)

type Connection struct {
	client *ent.Client
}

func New(config *Config) (*Connection, error) {
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Host,
			config.Port,
			config.User,
			config.DBName,
			config.Password,
		),
	)
	if err != nil {
		return nil, err
	}

	return &Connection{
		client: client,
	}, nil
}

func (c *Connection) OnStart(ctx context.Context) error {
	// Run the auto migration tool.
	err := c.client.Schema.Create(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) OnStop(ctx context.Context) error {
	err := c.client.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) GetClient() *ent.Client {
	return c.client
}

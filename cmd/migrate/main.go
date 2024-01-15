package main

import (
	"context"
	"flag"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"

	"github.com/Willsem/health-metrics/internal/config"
	"github.com/Willsem/health-metrics/internal/generated/ent/migrate"
	"github.com/Willsem/health-metrics/internal/startup"
)

const appName = "migrate"

func main() {
	ctx := context.Background()

	migrationName := flag.String("name", "", "name of current migration")
	flag.Parse()

	logger := startup.NewFallbackLogger(appName)

	if migrationName == nil || *migrationName == "" {
		logger.Fatal("name of migration could not be an empty")
	}

	config, err := config.NewMigration()
	if err != nil {
		logger.WithError(err).Fatal("create config")
	}

	dir, err := atlas.NewLocalDir("./internal/generated/migrations")
	if err != nil {
		logger.WithError(err).Fatal("create atlas migration directory")
	}

	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeInspect),
		schema.WithDialect(dialect.Postgres),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	url := config.Database.URL()
	err = migrate.NamedDiff(ctx, url, *migrationName, opts...)
	if err != nil {
		logger.WithError(err).Fatal("generating migration")
	}
}

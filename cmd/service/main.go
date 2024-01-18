package main

import (
	"context"

	"github.com/Willsem/health-metrics/internal/app"
	"github.com/Willsem/health-metrics/internal/app/build"
	"github.com/Willsem/health-metrics/internal/startup"
)

const appName = "health-metrics-api"

func main() {
	fallbackLogger := startup.NewFallbackLogger(appName)
	fallbackLogger.Infof(
		"%s has version %s built from %s on %s by %s",
		appName, build.Version, build.VersionCommit, build.BuildDate, build.GoVersion,
	)

	if err := app.Run(context.Background(), appName); err != nil {
		fallbackLogger.WithError(err).Fatal("failed to run the app")
	}
}

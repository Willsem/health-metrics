package main

import (
	"go.uber.org/fx"

	"github.com/Willsem/health-metrics/internal/app"
)

const appName = "health-metrics-api"

func main() {
	fx.New(
		app.WithLogger(),
		app.ProvideConfig(),
		app.ProvideLogger(appName),
		app.ProvideRouter(),
		app.ProvideHandlers(),
		app.RegisterHandlers(),
	).Run()
}

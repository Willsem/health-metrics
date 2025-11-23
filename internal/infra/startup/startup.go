package startup

import (
	"go.uber.org/fx"
)

type App struct {
	app *fx.App
}

func NewApp() *App {
	app := fx.New(
		// Infrastructure
		provideConfig(),
		provideLogger(),

		// Domain
		provideDomain(),

		// APIs
		provideBot(),

		// Usecases
		provideApplication(),

		// Repositories
		providePostgres(),
		provideContext(),
	)

	return &App{
		app: app,
	}
}

func (a *App) Run() {
	a.app.Run()
}

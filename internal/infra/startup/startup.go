package startup

import (
	"go.uber.org/fx"
)

type App struct {
	app *fx.App
}

func NewApp() *App {
	app := fx.New(
		provideConfig(),
		provideLogger(),
		provideBot(),
	)

	return &App{
		app: app,
	}
}

func (a *App) Run() {
	a.app.Run()
}

package startup

import (
	repositoryContextUser "health-metrics/internal/repository/context/user"

	"go.uber.org/fx"
)

func provideContext() fx.Option {
	return fx.Options(
		fx.Provide(repositoryContextUser.New),
	)
}

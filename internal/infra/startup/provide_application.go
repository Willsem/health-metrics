package startup

import (
	appGetOrgCreateUser "health-metrics/internal/application/getorcreateuser"
	repositoryPostgresUser "health-metrics/internal/repository/postgres/user"

	"go.uber.org/fx"
)

func provideApplication() fx.Option {
	return fx.Options(
		fx.Provide(
			func(repository *repositoryPostgresUser.UserRepository) *appGetOrgCreateUser.GetOrCreateUserUsecase {
				return appGetOrgCreateUser.New(repository)
			},
		),
	)
}

package startup

import (
	"health-metrics/internal/domain/converters"
	"health-metrics/internal/generated/ent"
	"health-metrics/internal/infra/postgres"
	userRepository "health-metrics/internal/repository/postgres/user"

	"go.uber.org/fx"
)

func providePostgres() fx.Option {
	return fx.Options(
		fx.Provide(postgres.New),

		fx.Invoke(func(lc fx.Lifecycle, connection *postgres.Connection) {
			lc.Append(fx.Hook{
				OnStart: connection.OnStart,
				OnStop:  connection.OnStop,
			})
		}),

		fx.Provide(func(connection *postgres.Connection) *ent.Client {
			return connection.GetClient()
		}),

		fx.Provide(
			func(client *ent.Client, userConverter *converters.User) *userRepository.UserRepository {
				return userRepository.New(client, userConverter)
			},
		),
	)
}

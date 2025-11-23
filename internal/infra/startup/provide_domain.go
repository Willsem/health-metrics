package startup

import (
	"health-metrics/internal/domain/converters"

	"go.uber.org/fx"
)

func provideDomain() fx.Option {
	return fx.Options(
		fx.Provide(converters.NewUserConverter),
	)
}

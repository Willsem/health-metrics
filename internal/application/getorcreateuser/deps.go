package getorcreateuser

import (
	"context"
	"health-metrics/internal/domain/model"
)

type UserRepository interface {
	CreateFromTelegramLogin(
		ctx context.Context,
		telegramLogin string,
	) (model.User, error)

	GetByTelegramLogin(
		ctx context.Context,
		telegramLogin string,
	) (model.User, error)
}

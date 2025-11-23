package auth

import (
	"context"
	domainModel "health-metrics/internal/domain/model"
)

type GetOrCreateUserUsecase interface {
	ByTelegramLogin(ctx context.Context, login string) (domainModel.User, error)
}

type UserContext interface {
	StoreUser(ctx context.Context, user domainModel.User) context.Context
}

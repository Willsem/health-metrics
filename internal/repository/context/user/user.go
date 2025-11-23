package user

import (
	"context"
	domainErrors "health-metrics/internal/domain/errors"
	domainModel "health-metrics/internal/domain/model"
)

type UserContext struct{}

func New() *UserContext {
	return &UserContext{}
}

func (c *UserContext) StoreUser(ctx context.Context, user domainModel.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func (c *UserContext) LoadUser(ctx context.Context) (domainModel.User, error) {
	user, ok := ctx.Value(userKey).(domainModel.User)
	if !ok {
		return domainModel.User{}, domainErrors.ErrNotFound
	}

	return user, nil
}

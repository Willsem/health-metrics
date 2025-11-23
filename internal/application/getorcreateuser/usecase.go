package getorcreateuser

import (
	"context"
	"errors"
	domainErrors "health-metrics/internal/domain/errors"
	"health-metrics/internal/domain/model"
)

type GetOrCreateUserUsecase struct {
	repository UserRepository
}

func New(repository UserRepository) *GetOrCreateUserUsecase {
	return &GetOrCreateUserUsecase{
		repository: repository,
	}
}

func (u *GetOrCreateUserUsecase) ByTelegramLogin(
	ctx context.Context,
	login string,
) (model.User, error) {
	user, err := u.repository.GetByTelegramLogin(ctx, login)
	if errors.Is(err, domainErrors.ErrNotFound) {
		user, err = u.repository.CreateFromTelegramLogin(ctx, login)
	}
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

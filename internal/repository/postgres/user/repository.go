package user

import (
	"context"
	"health-metrics/internal/domain/errors"
	"health-metrics/internal/domain/model"
	"health-metrics/internal/generated/ent"
	"health-metrics/internal/generated/ent/user"
)

type UserRepository struct {
	client    *ent.Client
	converter UserConverter
}

func New(client *ent.Client, converter UserConverter) *UserRepository {
	return &UserRepository{
		client:    client,
		converter: converter,
	}
}

func (r *UserRepository) CreateFromTelegramLogin(
	ctx context.Context,
	telegramLogin string,
) (model.User, error) {
	user, err := r.client.User.Create().
		SetTelegramLogin(telegramLogin).
		Save(ctx)
	if err != nil {
		return model.User{}, err
	}

	if user == nil {
		return model.User{}, nil
	}

	return r.converter.FromEntToDomain(*user), nil
}

func (r *UserRepository) GetByTelegramLogin(
	ctx context.Context,
	telegramLogin string,
) (model.User, error) {
	user, err := r.client.User.Query().
		Where(user.TelegramLogin(telegramLogin)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return model.User{}, errors.ErrNotFound
		}

		return model.User{}, err
	}

	if user == nil {
		return model.User{}, nil
	}

	return r.converter.FromEntToDomain(*user), nil
}

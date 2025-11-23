package user

import (
	"context"
	domainErrors "health-metrics/internal/domain/errors"
	domainModel "health-metrics/internal/domain/model"
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
) (domainModel.User, error) {
	user, err := r.client.User.Create().
		SetTelegramLogin(telegramLogin).
		Save(ctx)
	if err != nil {
		return domainModel.User{}, err
	}

	if user == nil {
		return domainModel.User{}, nil
	}

	return r.converter.FromEntToDomain(*user), nil
}

func (r *UserRepository) GetByTelegramLogin(
	ctx context.Context,
	telegramLogin string,
) (domainModel.User, error) {
	user, err := r.client.User.Query().
		Where(user.TelegramLogin(telegramLogin)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return domainModel.User{}, domainErrors.ErrNotFound
		}

		return domainModel.User{}, err
	}

	if user == nil {
		return domainModel.User{}, nil
	}

	return r.converter.FromEntToDomain(*user), nil
}

package converters

import (
	"health-metrics/internal/domain/model"
	"health-metrics/internal/generated/ent"
)

type User struct{}

func NewUserConverter() *User {
	return &User{}
}

func (*User) FromEntToDomain(entModel ent.User) model.User {
	domainModel := model.User{
		UUID:          entModel.ID,
		TelegramLogin: entModel.TelegramLogin,
	}

	if entModel.FatsecretAccessToken != "" {
		domainModel.FatsecretAccessToken = &entModel.FatsecretAccessToken
	}

	if entModel.FatsecretTokenSecret != "" {
		domainModel.FatsecretTokenSecret = &entModel.FatsecretTokenSecret
	}

	return domainModel
}

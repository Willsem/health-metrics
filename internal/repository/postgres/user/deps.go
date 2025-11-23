package user

import (
	"health-metrics/internal/domain/model"
	"health-metrics/internal/generated/ent"
)

type UserConverter interface {
	FromEntToDomain(entModel ent.User) model.User
}

package user

import (
	domainModel "health-metrics/internal/domain/model"
	"health-metrics/internal/generated/ent"
)

type UserConverter interface {
	FromEntToDomain(entModel ent.User) domainModel.User
}

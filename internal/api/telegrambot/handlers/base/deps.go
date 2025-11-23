package base

import (
	"context"
	domainModel "health-metrics/internal/domain/model"
)

type UserContext interface {
	LoadUser(ctx context.Context) (domainModel.User, error)
}

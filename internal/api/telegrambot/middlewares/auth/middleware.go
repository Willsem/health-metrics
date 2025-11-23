package auth

import (
	"context"
	"health-metrics/internal/infra/logger"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AuthMiddleware struct {
	getOrCreateUser GetOrCreateUserUsecase
	userContext     UserContext
	logger          logger.Logger
}

func New(
	getOrCreateUser GetOrCreateUserUsecase,
	userContext UserContext,
	logger logger.Logger,
) *AuthMiddleware {
	return &AuthMiddleware{
		getOrCreateUser: getOrCreateUser,
		userContext:     userContext,
		logger:          logger,
	}
}

func (m *AuthMiddleware) Middleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		user, err := m.getOrCreateUser.ByTelegramLogin(ctx, update.Message.From.Username)
		if err != nil {
			m.logger.WithError(err).Error("не удалось получить пользователя")
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "⚠️ Ошибка авторизации, обратитесь к владельцу бота за помощью",
			})
			return
		}

		ctx = m.userContext.StoreUser(ctx, user)
		next(ctx, b, update)
	}
}

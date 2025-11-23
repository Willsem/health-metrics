package base

import (
	"context"
	"encoding/json"
	"health-metrics/internal/infra/logger"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handler struct {
	userContext UserContext
	logger      logger.Logger
}

func New(userContext UserContext, logger logger.Logger) *Handler {
	return &Handler{
		userContext: userContext,
		logger:      logger,
	}
}

func (h *Handler) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.userContext.LoadUser(ctx)
	if err != nil {
		h.logger.WithError(err).Error("не удалось получить авторизованного пользователя")

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "⚠️ Ошибка авторизации, обратитесь к владельцу бота за помощью",
		})
		if err != nil {
			h.logger.WithError(err).Error("не удалось отправить сообщение")
		}
	}

	data, _ := json.MarshalIndent(user, "  ", "")

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   string(data),
	})
	if err != nil {
		h.logger.WithError(err).Error("не удалось отправить сообщение")
	}
}

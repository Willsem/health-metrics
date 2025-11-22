package base

import (
	"context"
	"health-metrics/internal/infra/logger"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handler struct {
	logger logger.Logger
}

func New(logger logger.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "base_handler",
	})
	if err != nil {
		h.logger.WithError(err).Error("failed to send the message")
	}
}

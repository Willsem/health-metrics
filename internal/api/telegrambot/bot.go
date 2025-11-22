package telegrambot

import "context"

type Bot struct{}

func New() *Bot {
	return &Bot{}
}

func (b *Bot) Start(ctx context.Context) error {
	return nil
}

func (b *Bot) Shutdown(ctx context.Context) error {
	return nil
}

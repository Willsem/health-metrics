package telegrambot

import (
	"context"
	"health-metrics/internal/api/telegrambot/handlers/base"

	"github.com/go-telegram/bot"
)

type Bot struct {
	bot  *bot.Bot
	stop chan struct{}
	done chan struct{}
}

func New(
	config *Config,
	baseHandler *base.Handler,
) (*Bot, error) {
	bot, err := bot.New(
		config.Token,
		bot.WithDefaultHandler(baseHandler.Handle),
	)
	if err != nil {
		return nil, err
	}

	return &Bot{
		bot:  bot,
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}, nil
}

func (b *Bot) Start(ctx context.Context) error {
	go b.run()
	return nil
}

func (b *Bot) Shutdown(ctx context.Context) error {
	close(b.stop)

	select {
	case <-b.done:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func (b *Bot) run() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		b.bot.Start(ctx)
		close(b.done)
	}()

	<-b.stop
	cancel()
}

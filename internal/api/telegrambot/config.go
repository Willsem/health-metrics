package telegrambot

type Config struct {
	Token string `envconfig:"TOKEN" required:"true"`
}

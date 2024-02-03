package router

type Config struct {
	Port int `envconfig:"PORT" required:"true"`
}

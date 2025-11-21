package logger

type Config struct {
	Env  string `envconfig:"ENV"  default:"local"`
	Name string `envconfig:"NAME" default:"health-metrics"`
}

package postgres

type Config struct {
	Host     string `envconfig:"HOST"     required:"true"`
	Port     int    `envconfig:"PORT"     required:"true"`
	User     string `envconfig:"USER"     required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DBName   string `envconfig:"DB_NAME"  required:"true"`
}

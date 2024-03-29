package database

import "fmt"

type Config struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     int    `envconfig:"PORT" required:"true"`
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`
	SSlMode  string `envconfig:"SSL_MODE" default:"disable"`
}

const (
	connectionStringFormat = "host=%s port=%d user=%s dbname=%s password=%s sslmode=%s"
	urlFormat              = "postgresql://%s:%s@%s:%d/%s?sslmode=%s"
)

func (c *Config) ConnectionString() string {
	return fmt.Sprintf(connectionStringFormat, c.Host, c.Port, c.User, c.DBName, c.Password, c.SSlMode)
}

func (c *Config) URL() string {
	return fmt.Sprintf(urlFormat, c.User, c.Password, c.Host, c.Port, c.DBName, c.SSlMode)
}

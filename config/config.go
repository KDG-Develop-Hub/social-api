package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type (
	Config struct {
		DB Database
	}
	Database struct {
		Host     string `env:"DB_HOST" env-required:"true"`
		Port     int    `env:"DB_PORT" env-required:"true"`
		User     string `env:"DB_USER" env-required:"true"`
		Password string `env:"DB_PASSWORD" env-required:"true"`
		Name     string `env:"DB_NAME" env-required:"true"`
		SslMode  string `env:"DB_SSL_MODE" env-required:"true"`
	}
)

func MustLoad() *Config {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &cfg
}

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
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		Port     int    `env:"DB_PORT" envDefault:"5432"`
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASSWORD" envDefault:"postgres"`
		Name     string `env:"DB_NAME" envDefault:"postgres"`
		SslMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
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

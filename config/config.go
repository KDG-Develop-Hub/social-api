package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		DB Database `env-prefix:"DB_"`
	}
	Database struct {
		Host     string `env:"HOST" env-required:"true"`
		Port     int    `env:"PORT" env-required:"true"`
		User     string `env:"USER" env-required:"true"`
		Password string `env:"PASSWORD" env-required:"true"`
		Name     string `env:"NAME" env-required:"true"`
		SslMode  string `env:"SSL_MODE" env-required:"true"`
	}
)

func MustLoad() *Config {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &cfg
}

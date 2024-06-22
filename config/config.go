package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port   string
	Env    string
	Secret string
}

type DbConfig struct {
	Url string
}

type Config struct {
	App AppConfig
	Db  DbConfig
}

func LoadConfig() (*Config, error) {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	appConfig := AppConfig{
		Port:   os.Getenv("APP_PORT"),
		Env:    os.Getenv("APP_ENV"),
		Secret: os.Getenv("APP_SECRET"),
	}

	dbConfig := DbConfig{
		Url: os.Getenv("DB_URL"),
	}

	return &Config{
		App: appConfig,
		Db:  dbConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}

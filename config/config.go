package config

import (
	"os"
	"strconv"

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

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type JwtConfig struct {
	Secret        string
	AccessTTL     int
	RefreshTTL    int
	Issuer        string
	ResetTokenTTL int
}

type Config struct {
	App   AppConfig
	Db    DbConfig
	Redis RedisConfig
	Jwt   JwtConfig
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

	redisPort, err := strconv.ParseInt(os.Getenv("REDIS_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	redisConfig := RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     int(redisPort),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	accessTTL, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_TTL"), 10, 64)
	if err != nil {
		return nil, err
	}
	refreshTTL, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_TTL"), 10, 64)
	if err != nil {
		return nil, err
	}
	resetTokenTTL, err := strconv.ParseInt(os.Getenv("JWT_RESET_TOKEN_TTL"), 10, 64)
	if err != nil {
		return nil, err
	}

	jwtConfig := JwtConfig{
		Secret:        os.Getenv("JWT_SECRET"),
		AccessTTL:     int(accessTTL),
		RefreshTTL:    int(refreshTTL),
		Issuer:        os.Getenv("JWT_ISSUER"),
		ResetTokenTTL: int(resetTokenTTL),
	}

	return &Config{
		App:   appConfig,
		Db:    dbConfig,
		Redis: redisConfig,
		Jwt:   jwtConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}

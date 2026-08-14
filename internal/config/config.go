package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
	MongoDB  MongoConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}
type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (p PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Jakarta", p.Host, p.Port, p.Username, p.Database, p.Password)
}

type JWTConfig struct {
	Secret             string
	ExpiryHours        time.Duration
	RefreshExpiryHours time.Duration
}

type MongoConfig struct {
	URI    string
	DBName string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtExpiry, err := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRY_HOURS: %w", err)
	}

	jwtRefreshExpiry, err := strconv.Atoi(getEnv("JWT_REFRESH_EXPIRY_HOURS", "168"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_REFRESH_EXPIRY_HOURS: %w", err)
	}

	return &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "e-massagger"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Postgres: PostgresConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			Username: getEnv("POSTGRES_USER", "appuser"),
			Password: getEnv("POSTGRES_PASSWORD", "secretpassword"),
			Database: getEnv("POSTGRES_DB", "messengerdb"),
		},
		MongoDB: MongoConfig{
			URI:    getEnv("MONGO_URI", "mongodb://mongoadmin:secretpassword@localhost:27017"),
			DBName: getEnv("MONGO_DB", "emassagger"),
		},
		JWT: JWTConfig{
			Secret:             getEnv("JWT_SECRET", "secret"),
			ExpiryHours:        time.Duration(jwtExpiry) * time.Hour,
			RefreshExpiryHours: time.Duration(jwtRefreshExpiry) * time.Hour,
		},
	}, nil
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string
	DBHost        string
	DBName        string
	DBUser        string
	DBPassword    string
	DBPort        string
	SessionSecret string
}

func LoadConfig() (*Config, error) {
	godotenv.Load()

	return &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBName:        getEnv("DB_NAME", "edukarsa"),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", ""),
		DBPort:        getEnv("DB_PORT", "5432"),
		SessionSecret: getEnv("SESSION_SECRET", "abcdefg"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	return defaultValue
}

package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	DatabasePath  string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "https://tire-shop.zeabur.app:8080"),
		DatabasePath:  getEnv("DATABASE_PATH", "./database.db"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return fallback
}

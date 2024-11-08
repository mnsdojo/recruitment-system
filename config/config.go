package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost     string
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func NewConfig() *Config {
	return &Config{
		Port:       getEnv("PORT", "8000"),
		DBHost:     getEnv("DB_HOST", "go_db"),
		DBUser:     getEnv("DB_USER", "your_db_user"),
		DBPassword: getEnv("DB_PASSWORD", "your_password"),
		DBName:     getEnv("DB_NAME", "your_db_name"),
		DBPort:     getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBPort,
	)
}

func (c *Config) GetPort() string {
	return c.Port
}

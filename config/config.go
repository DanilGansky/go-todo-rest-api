package config

import (
	"os"
	"sync"
)

// Config ...
type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	AppHost    string
	AppPort    string
}

var (
	configInstance *Config
	once           sync.Once
)

// GetConfig returns config
func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{
			DBName:     os.Getenv("DB_NAME"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			AppHost:    os.Getenv("APP_HOST"),
			AppPort:    os.Getenv("APP_PORT"),
		}
	})
	return configInstance
}

package config

import "os"

type Config struct {
	DB_CONN   string
	APP_PORT  string
	HOST_PORT string
}

func NewConfig() *Config {
	return &Config{
		DB_CONN:   os.Getenv("DB_CONN"),
		APP_PORT:  os.Getenv("APP_PORT"),
		HOST_PORT: os.Getenv("HOST_PORT"),
	}
}

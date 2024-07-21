package config

import "os"

type Config struct {
	DB_CONN string
}

func NewConfig() *Config {
	return &Config{
		DB_CONN: os.Getenv("DB_CONN"),
	}
}

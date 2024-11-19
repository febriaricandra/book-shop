package config

import (
	"os"
)

type Config struct {
	JWTSecret []byte
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret: []byte(os.Getenv("JWT_SECRET")),
	}
}

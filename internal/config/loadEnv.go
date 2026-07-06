package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	PG_DSN    string
	SecretKey string
}

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	confgoenv := &Config{
		Port:      os.Getenv("PORT"),
		PG_DSN:    os.Getenv("PG_DSN"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	return confgoenv
}

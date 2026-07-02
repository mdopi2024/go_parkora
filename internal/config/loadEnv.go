package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	PG_DSN string
}

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	confgoenv := &Config{
		Port:   os.Getenv("PORT"),
		PG_DSN: os.Getenv("PG_DSN"),
	}

	return confgoenv
}

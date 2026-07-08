// package config

// import (
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type Config struct {
// 	Port      string
// 	PG_DSN    string
// 	SecretKey string
// }

// func LoadEnv() *Config {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic(err)
// 	}
// 	confgoenv := &Config{
// 		Port:      os.Getenv("PORT"),
// 		PG_DSN:    os.Getenv("PG_DSN"),
// 		SecretKey: os.Getenv("SECRET_KEY"),
// 	}

// 	return confgoenv
// }

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	PG_DSN    string
	SecretKey string
}

func LoadEnv() *Config {
	// Try to load .env file (for local development only)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Port:      os.Getenv("PORT"),
		PG_DSN:    os.Getenv("PG_DSN"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}

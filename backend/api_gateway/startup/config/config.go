package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	AuthenticationHost string
	AuthenticationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:               os.Getenv("GATEWAY_PORT"),
		AuthenticationHost: os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort: os.Getenv("AUTHENTICATION_SERVICE_PORT"),
	}
}

func SetEnvironment() error {
	if os.Getenv("OS_ENV") != "docker" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return nil
}

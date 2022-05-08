package config

import (
	"os"
)

type Config struct {
	Port               string
	AuthenticationHost string
	AuthenticationPort string
	ProfileHost        string
	ProfilePort        string
}

func NewConfig() *Config {
	return &Config{
		Port:               os.Getenv("GATEWAY_PORT"),
		AuthenticationHost: os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort: os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		ProfileHost:        os.Getenv("PROFILE_SERVICE_HOST"),
		ProfilePort:        os.Getenv("PROFILE_SERVICE_PORT"),
	}
}

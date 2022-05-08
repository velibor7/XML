package config

import "os"

type Config struct {
	Port                 string
	AuthenticationDBHost string
	AuthenticationDBPort string
	NatsHost             string
	NatsPort             string
	NatsUser             string
	NatsPass             string
}

func NewConfig() *Config {
	return &Config{
		Port:                 os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		AuthenticationDBHost: os.Getenv("AUTHENTICATION_DB_HOST"),
		AuthenticationDBPort: os.Getenv("AUTHENTICATION_DB_PORT"),
		NatsHost:             os.Getenv("NATS_HOST"),
		NatsPort:             os.Getenv("NATS_PORT"),
		NatsUser:             os.Getenv("NATS_USER"),
		NatsPass:             os.Getenv("NATS_PASS"),
	}
}

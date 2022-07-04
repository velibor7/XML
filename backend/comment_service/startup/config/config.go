package config

import "os"

type Config struct {
	Port                        string
	CommentDBHost               string
	CommentDBPort               string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	UpdateProfileCommandSubject string
	UpdateProfileReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                        os.Getenv("COMMENT_SERVICE_PORT"),
		CommentDBHost:               os.Getenv("COMMENT_DB_HOST"),
		CommentDBPort:               os.Getenv("COMMENT_DB_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
	}
}

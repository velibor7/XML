package config

import (
	"os"
)

type Config struct {
	Port                        string
	PostDBHost                  string
	PostDBPort                  string
	NatsHost                    string
	NatsPort                    string
	NatsUser                    string
	NatsPass                    string
	UpdateProfileCommandSubject string
	UpdateProfileReplySubject   string
	CommentHost                 string
	CommentPort                 string
}

func NewConfig() *Config {
	return &Config{
		Port:                        os.Getenv("POST_SERVICE_PORT"),
		PostDBHost:                  os.Getenv("POST_DB_HOST"),
		PostDBPort:                  os.Getenv("POST_DB_PORT"),
		CommentHost:                 os.Getenv("COMMENT_SERVICE_HOST"),
		CommentPort:                 os.Getenv("COMMENT_SERVICE_PORT"),
		NatsHost:                    os.Getenv("NATS_HOST"),
		NatsPort:                    os.Getenv("NATS_PORT"),
		NatsUser:                    os.Getenv("NATS_USER"),
		NatsPass:                    os.Getenv("NATS_PASS"),
		UpdateProfileCommandSubject: os.Getenv("UPDATE_PROFILE_COMMAND_SUBJECT"),
		UpdateProfileReplySubject:   os.Getenv("UPDATE_PROFILE_REPLY_SUBJECT"),
	}
}

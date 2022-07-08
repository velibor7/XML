package main

import (
	"github.com/velibor7/XML/authentication_service/startup"
	"github.com/velibor7/XML/authentication_service/startup/config"
	"github.com/velibor7/XML/common/loggers"
)

var log = loggers.NewAuthenticationLogger()

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	log.Info("Authentication service started! \n")
	server.Start()
}

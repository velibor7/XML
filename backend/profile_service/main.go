package main

import (
	"github.com/velibor7/XML/common/loggers"
	"github.com/velibor7/XML/profile_service/startup"
	"github.com/velibor7/XML/profile_service/startup/config"
)

var log = loggers.NewProfileLogger()

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	log.Info("Profile service started!\n")
	server.Start()
}

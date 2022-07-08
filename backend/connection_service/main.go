package main

import (
	"github.com/velibor7/XML/common/loggers"
	"github.com/velibor7/XML/connection_service/startup"
	cfg "github.com/velibor7/XML/connection_service/startup/config"
)

var log = loggers.NewConnectionLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Connection service started\n")
	server.Start()
}

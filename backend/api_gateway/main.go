package main

import (
	"github.com/velibor7/XML/api_gateway/startup"
	"github.com/velibor7/XML/api_gateway/startup/config"
	"github.com/velibor7/XML/common/loggers"
)

var log = loggers.NewGatewayLogger()

func main() {

	config := config.NewConfig()
	server := startup.NewServer(config)
	log.Info("API Gateway started")
	server.Start()
}

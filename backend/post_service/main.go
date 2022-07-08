package main

import (
	"github.com/velibor7/XML/common/loggers"
	"github.com/velibor7/XML/post_service/startup"
	cfg "github.com/velibor7/XML/post_service/startup/config"
)

var log = loggers.NewPostLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Post service started \n")
	server.Start()
}

package main

import (
	"github.com/velibor7/XML/comment_service/startup"
	cfg "github.com/velibor7/XML/comment_service/startup/config"
	"github.com/velibor7/XML/common/loggers"
)

var log = loggers.NewJobLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Comment service startted \n")
	server.Start()

}

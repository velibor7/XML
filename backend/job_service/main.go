package main

import (
	"github.com/velibor7/XML/common/loggers"
	"github.com/velibor7/XML/job_service/startup"
	cfg "github.com/velibor7/XML/job_service/startup/config"
)

var log = loggers.NewJobLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Job service started\n")
	server.Start()

}

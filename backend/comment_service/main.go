package main

import (
	cfg "github.com/velibor7/XML/comment_service/startup/config"

	"github.com/velibor7/XML/comment_service/startup"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()

}

// pokrenuti nats server negde

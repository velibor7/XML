package main

import (
	"fmt"

	"github.com/velibor7/XML/post_service/startup"
	cfg "github.com/velibor7/XML/post_service/startup/config"
)

func main() {
	fmt.Println("Server started.")
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

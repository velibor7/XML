package main

import (
	"github.com/velibor7/XML/profile_service/startup"
	"github.com/velibor7/XML/profile_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()

	// http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(":2112", nil)
}

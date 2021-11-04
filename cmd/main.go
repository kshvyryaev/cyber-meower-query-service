package main

import (
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/di"
)

func main() {
	server, cleanup, err := di.InitializeHttpServer()
	if err != nil {
		panic("cannot initialize http server: " + err.Error())
	}
	defer cleanup()

	server.Run()
}

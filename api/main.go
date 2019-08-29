package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"xiaoshijie.com/micro/hello/api/client"
	"xiaoshijie.com/micro/hello/api/handler"

	api "xiaoshijie.com/micro/hello/api/proto/api"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.apiHelloService"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Api srv client
		micro.WrapHandler(client.ApiWrapper(service)),
	)

	// Register Handler
	api.RegisterApiHandler(service.Server(), new(handler.Api))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

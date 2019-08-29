package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client/grpc"
	_ "github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/util/log"
	"xiaoshijie.com/micro/hello/api/client"
	"xiaoshijie.com/micro/hello/api/handler"
	
	api "xiaoshijie.com/micro/hello/api/proto/api"
)
func GrpcClient(v string) micro.Option {
	return func(o *micro.Options) {
		o.Client = grpc.NewClient()
	}
}
func main() {
	
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.apiHelloService"),
		micro.Version("latest"),
		GrpcClient("1"),
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

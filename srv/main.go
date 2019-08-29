package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"xiaoshijie.com/micro/hello/srv/handler"
	"xiaoshijie.com/micro/hello/srv/subscriber"
	
	"github.com/micro/go-micro/service/grpc"
	
	helloService "xiaoshijie.com/micro/hello/srv/proto/helloService"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.helloService"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	helloService.RegisterHelloServiceHandler(service.Server(), new(handler.HelloService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.helloService", service.Server(), new(subscriber.HelloService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.helloService", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

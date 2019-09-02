package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"xiaoshijie.com/micro/hello/srv/handler"
	helloService "../srv/proto/helloService"
	"xiaoshijie.com/micro/hello/srv/provider"
	"xiaoshijie.com/micro/hello/srv/subscriber"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.helloService"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(micro.Action(func(context *cli.Context) {
		log.Logf("start application")
		// Create new config
		conf := config.NewConfig()
		conf.Load(provider.EtcdSource)
	}),micro.BeforeStart(func() error {
		
		return nil
	}),micro.AfterStop(func() error {
		return nil
	}))

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

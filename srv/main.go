package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"xiaoshijie.com/micro/common/bootstrap"
	provider2 "xiaoshijie.com/micro/common/provider"
	"xiaoshijie.com/micro/hello/srv/config"
	"xiaoshijie.com/micro/hello/srv/handler"
	helloService "xiaoshijie.com/micro/hello/srv/proto/helloService"
	"xiaoshijie.com/micro/hello/srv/provider"
	"xiaoshijie.com/micro/hello/srv/subscriber"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.helloService"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:  "environment",
				Usage: "The environment",
			},
		),
	)

	// Initialise service
	service.Init(micro.Action(func(context *cli.Context) {
		log.Logf("start application")
		env := context.String("environment")
		if len(env) > 0 {
			fmt.Println("Environment set to", env)
		}
		bootstrap.App.Register(provider2.NewProviderEtcd("127.0.0.1:2379","/micro/config/develop", true))
		bootstrap.App.Register(provider.NewProviderConfig(&config.ServiceConf, env))
		bootstrap.App.Register(provider.NewProviderRedis(&config.ServiceConf))
		bootstrap.App.Register(provider.NewProviderMysql(&config.ServiceConf))
		bootstrap.App.Register(provider.NewProviderService(&config.ServiceConf))
	
		bootstrap.App.Boot()
		
		
		
	}),micro.BeforeStart(func() error {
		bootstrap.App.After()
		return nil
	}),micro.AfterStop(func() error {
		bootstrap.App.Shutdown()
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

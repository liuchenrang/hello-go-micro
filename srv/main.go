package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"xiaoshijie.com/micro/common/bootstrap"
	provider2 "xiaoshijie.com/micro/common/provider"
	config2 "xiaoshijie.com/micro/hello/srv/config"
	"xiaoshijie.com/micro/hello/srv/handler"
	helloService "xiaoshijie.com/micro/hello/srv/proto/helloService"
	"xiaoshijie.com/micro/hello/srv/subscriber"
	
	_ "github.com/micro/go-plugins/registry/etcdv3"
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
		bootstrap.App.Register(provider2.NewProviderEtcd("127.0.0.1:2379","/micro/config/develop", true))
		bootstrap.App.Register(provider2.NewProviderRedis(&config2.ServiceConf))
	
		bootstrap.App.Boot()
		env := context.String("environment")
		if len(env) > 0 {
			fmt.Println("Environment set to", env)
		}
		if provider2.EtcdSource != nil {
			err := config.Load(provider2.EtcdSource)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			config.Get("hello").Scan(&config2.ServiceConf)
			conf2 := config.Map()
			fmt.Printf("%+v",conf2)
		}else{
			println("etcd init fail")
		}
		
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

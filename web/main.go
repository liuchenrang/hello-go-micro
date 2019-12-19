package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/util/log"
	"xiaoshijie.com/micro/common/bootstrap"
	provider2 "xiaoshijie.com/micro/common/provider"
	"xiaoshijie.com/micro/tbservice/tb-service-srv/config"
	"xiaoshijie.com/micro/tbservice/tb-service-srv/provider"
	
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"net/http"
	"os"
	"time"
	"xiaoshijie.com/micro/tbservice/tb-service-web/router"
)


func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.tb-service-web"),
		web.Version("latest"),
		web.RegisterInterval(time.Second*5),
		web.RegisterTTL(time.Second*10),
	)
	
	// initialise service
	var err error
	
	err = service.Init(web.Action(func(context *cli.Context) {
		env := os.Getenv("APPLICATION_ENV")
		if env == "" {
			env = "production"
		}
		username := os.Getenv("REGISTER_AUTH_USERNAME")
		password := os.Getenv("REGISTER_AUTH_PASSWORD")
		options := service.Options()
		err := options.Service.Options().Registry.Init(etcdv3.Auth(username, password))
		
		client.DefaultClient = grpc.NewClient(func(ops *client.Options) {
			ops.Registry = options.Service.Options().Registry
		})
		if err != nil {
			panic(err.Error())
		}
		bootstrap.App.Register(provider2.NewProviderEtcd(context.String("registry_address"),"/micro/config/"+env, true, username, password))
		bootstrap.App.Register(provider.NewProviderConfig(&config.ServiceConf, env))
		bootstrap.App.Register(provider.NewProviderLogger(&config.ServiceConf, env))
		bootstrap.App.Register(provider.NewProviderRedis(&config.ServiceConf))
		bootstrap.App.Register(provider.NewProviderMysql(&config.ServiceConf))
		bootstrap.App.Register(provider.NewProviderService(&config.ServiceConf))
		bootstrap.App.Register(provider.NewProviderQueue(&config.ServiceConf))
		bootstrap.App.Register(provider.NewProviderCounter(&config.ServiceConf))
		bootstrap.App.Boot()
		bootstrap.App.After()
	}),web.BeforeStop(func() error {
		bootstrap.App.Shutdown()
		return nil
	}))
	
	if  err != nil {
		log.Error(err.Error())
		return
	}

	// register html handler
	service.Handle("/static", http.FileServer(http.Dir("html")))
	service.Handle("/", router.GetRouter())
	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

package provider

import (
	config2 "github.com/micro/go-micro/config"
	"log"
	"duoduo.com/micro/common/provider"
	"duoduo.com/micro/hello/srv/config"
)

type ProviderConfig struct {
	config *config.ServiceConfig
	env string
}

func (c *ProviderConfig) Boot()  {
	print("boot config")
	
	if provider.EtcdSource != nil {
		err := config2.Load(provider.EtcdSource)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		err = config2.Get("hello").Scan(&config.ServiceConf)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}else{
		println("etcd init fail")
	}
}
func (c *ProviderConfig) Shutdown()  {

}
func (c *ProviderConfig) After()  {

}

func NewProviderConfig(config *config.ServiceConfig,env string) *ProviderConfig {
	return &ProviderConfig{config: config,env:env}
}
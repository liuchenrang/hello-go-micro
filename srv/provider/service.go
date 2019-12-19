package provider

import (
	"duoduo.com/micro/hello/srv/config"
	"duoduo.com/micro/hello/srv/services"
)

var (

)

func init() {
	
}

type ProviderService struct {
	config *config.ServiceConfig
}

func NewProviderService(config *config.ServiceConfig) *ProviderService {
	return &ProviderService{config: config}
}
func (e *ProviderService) Boot() {
	println("boot service")
	
}
func (e *ProviderService) After() {
	services.SqbApp = services.NewServiceSqbApp(&config.ServiceConf)
	services.SqbUser = services.NewServiceSqbUser(&config.ServiceConf)
}
func (e *ProviderService) Shutdown() {

}

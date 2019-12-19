package provider

import (
	"github.com/go-redis/redis"
	"time"
	config2 "duoduo.com/micro/common/config"
	"duoduo.com/micro/hello/srv/config"
)

var ()

func init() {

}

type ProviderRedis struct {
	config *config.ServiceConfig
}

func NewProviderRedis(config *config.ServiceConfig) *ProviderRedis {
	return &ProviderRedis{config: config}
}
func (e *ProviderRedis) Boot() {
	print("boot redis")
}
func (e *ProviderRedis) After() {
	for key, redisConf := range e.config.RedisGroup {
		client, err := e.createClient(redisConf)
		if err != nil {
			config.ServiceConf.ConnectGroup.RedisGroup[key] = client
		}
	}
}
func (e *ProviderRedis) Shutdown() {
	for _,connect :=range 			config.ServiceConf.ConnectGroup.RedisGroup{
		connect.(*redis.Client).Close()
	}
}
func (e *ProviderRedis) createClient(conf config2.RedisConf) (*redis.Client, error){
	var poolSize int
	if conf.PoolSzie > 0 {
		poolSize = conf.PoolSzie
	}else{
		poolSize = 10
	}
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password, // no password set
		DB:       0,             // use default DB
		PoolSize: poolSize,
		IdleCheckFrequency: time.Second * 5,
	})
	
	return client, nil
}

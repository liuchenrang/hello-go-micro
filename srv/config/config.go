package config

import "duoduo.com/micro/common/config"

var (
	ServiceConf ServiceConfig
)

func init()  {
	ServiceConf.MysqlGroup = make(map[string]config.MysqlConf)
	ServiceConf.RedisGroup = make(map[string]config.RedisConf)
	ServiceConf.ConnectGroup = config.NewConnectGroup()
}
type ServiceConfig struct {
	RedisGroup map[string]config.RedisConf `json:"redis"`
	MysqlGroup map[string]config.MysqlConf `json:"mysql"`
	
	ConnectGroup *config.ConnectGroup
}


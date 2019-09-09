package config

import "xiaoshijie.com/micro/common/config"

var (
	ServiceConf ServiceConfig
)
type ServiceConfig struct {
	RedisGroup map[string]config.RedisConf `json:"redis"`
	MysqlGroup map[string]config.MysqlConf `json:"mysql"`
	
	ConnectGroup config.ConnectGroup
}


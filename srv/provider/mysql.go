package provider

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"duoduo.com/go-xorm/xorm"
	config2 "duoduo.com/micro/common/config"
	"duoduo.com/micro/hello/srv/config"
)

var ()

func init() {

}

type ProviderMysql struct {
	config *config.ServiceConfig
}

func NewProviderMysql(config *config.ServiceConfig) *ProviderMysql {
	return &ProviderMysql{config: config}
}
func (e *ProviderMysql) Boot() {
	println("boot mysql")
	for key, redisConf := range e.config.MysqlGroup {
		client, err := e.createClient(redisConf)
		if err == nil {
			config.ServiceConf.ConnectGroup.MysqlGroup[key] = client
		}
	}

}
func (e *ProviderMysql) After() {

}
func (e *ProviderMysql) Shutdown() {
	for _, connect := range config.ServiceConf.ConnectGroup.RedisGroup {
		connect.(*xorm.Engine).Close()
	}
}
func (e *ProviderMysql) createClient(conf config2.MysqlConf) (*xorm.Engine, error) {
	var idle, open, port int
	if conf.MaxIdleConns > 0 {
		idle = conf.MaxIdleConns
	} else {
		idle = 10
	}
	if conf.MaxOpenConns > 0 {
		open = conf.MaxOpenConns
	} else {
		open = 16
	}
	port = 3306
	if conf.Port > 0 {
		port = conf.Port
	}
	//database: xsj_test_hd:t2t_e0e_s1s_t5t@tcp(rm-bp1eifgih690z1h372o.mysql.rds.aliyuncs.com:3306)/xsj_master?charset=utf8&parseTime=True&loc=Local
	//database:xsj_test_hd:t2t_e0e_s1s_t5t@tcp(rm-bp1eifgih690z1h372o.mysql.rds.aliyuncs.com:3306)/?charset=utf8&parseTime=True&loc=Local
	//database:xsj_test_hd:t2t_e0e_s1s_t5t@tcp(rm-bp1eifgih690z1h372o.mysql.rds.aliyuncs.com:3306)/xsj_master?charset=utf8&parseTime=True&loc=Local
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.UserName,
		conf.PassWord,
		conf.Addr,
		port,
		conf.Database,
	)
	eg, err := xorm.NewEngine("mysql", dns)
	if err != nil {
		panic("failed to connect database " + err.Error())
	}
	eg.ShowSQL(true)
	
	eg.SetSoftDeleteHandler(&xorm.DefaultSoftDeleteHandler{})
	
	eg.SetMaxIdleConns(idle)
	eg.SetMaxOpenConns(open)
	
	return eg, nil
}

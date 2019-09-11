package services

import (
	"xiaoshijie.com/go-xorm/xorm"
	"xiaoshijie.com/micro/common/components"
	"xiaoshijie.com/micro/hello/srv/contract/daos"
	"xiaoshijie.com/micro/hello/srv/contract/services"
	
	"xiaoshijie.com/micro/hello/srv/config"
	
	"xiaoshijie.com/micro/hello/srv/daos"
	
	"xiaoshijie.com/micro/hello/srv/models"
)

var (
	SqbApp contractService.IServiceSqbApp
)

type ServiceSqbApp struct {
	dao  contractDao.ISqbAppDao
	conf *config.ServiceConfig
}

func (c ServiceSqbApp) Add(object models.SqbApp) (models.SqbApp, error) {
	return c.dao.Create(object)
}
func (c ServiceSqbApp) Delete(object models.SqbApp) (int, error) {
	return c.dao.Delete(object)
}
func (c ServiceSqbApp) DeleteById(id uint) (int, error) {
	return c.dao.DeleteById(id)
}

func (c ServiceSqbApp) Update(id uint, attr map[string]interface{}) (bool, error) {
	return c.dao.Update(id, attr)
}
func (c ServiceSqbApp) UpdateByModel(model models.SqbApp) (bool, error) {
	return c.dao.UpdateByModel(model)
}

func (c ServiceSqbApp) Find(id uint) (*models.SqbApp, error) {
	return c.dao.Find(id)
}
func (c ServiceSqbApp) FindCode(code string) (*models.SqbApp, error) {
	if c.dao == nil {
		print("dao nill")
	} else {
		print("dao inistance")
	}
	return c.dao.FindByKey("code", code)
}

func (c ServiceSqbApp) FindAll(page components.Page) (*[]*models.SqbApp, error) {
	return c.dao.FindAll(page)
}
func (c ServiceSqbApp) FindWithID(ids []uint) (*[]*models.SqbApp, error) {
	return c.dao.FindWithID(ids)
}

func (c ServiceSqbApp) CountWhereForAdmin(wh string, bind ...string) (int64, error) {
	return c.dao.CountWhereForAdmin(wh, bind...)
}

func (c ServiceSqbApp) FindWhereForAdmin(wh string, page components.Page, bind ...string) (*[]*models.SqbApp, error) {
	return c.dao.FindWhereForAdmin(wh, page, bind...)
}
func (c ServiceSqbApp) Count() (int64, error) {
	return c.dao.Count()
}

func NewServiceSqbApp(conf *config.ServiceConfig) ServiceSqbApp {
	engine := conf.ConnectGroup.MysqlGroup["site_master"].(*xorm.Engine)
	
	dao := daos.NewSqbAppDao(conf, engine)
	if dao == nil {
		print("dao nill")
	} else {
		print("dao inistance")
	}
	return ServiceSqbApp{conf: conf, dao: dao}
}

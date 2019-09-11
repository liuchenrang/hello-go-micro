package services

import (
 
        "xiaoshijie.com/go-xorm/xorm"
 
        "xiaoshijie.com/micro/common/components"
 
        contractDao "xiaoshijie.com/micro/hello/srv/contract/daos"
 
        contractService "xiaoshijie.com/micro/hello/srv/contract/services"
 
        "xiaoshijie.com/micro/hello/srv/config"
 
        "xiaoshijie.com/micro/hello/srv/daos"
 
        "xiaoshijie.com/micro/hello/srv/models"
 
)

var (
    SqbUser contractService.IServiceSqbUser
)

type ServiceSqbUser struct {
    dao contractDao.ISqbUserDao
    conf *config.ServiceConfig

}

func (c ServiceSqbUser) Add(object models.SqbUser)  (models.SqbUser, error) {
	return c.dao.Create(object)
}
func (c ServiceSqbUser) Delete(object models.SqbUser) (int, error){
	return c.dao.Delete(object)
}
func (c ServiceSqbUser) DeleteById(id uint) (int, error) {
	return c.dao.DeleteById(id)
}



func (c ServiceSqbUser) Update(id uint,  attr map[string]interface{})  (bool, error) {
	return c.dao.Update(id, attr)
}
func (c ServiceSqbUser) UpdateByModel(model models.SqbUser)  (bool, error){
	return c.dao.UpdateByModel(model)
}

func (c ServiceSqbUser) Find(id uint) ( *models.SqbUser, error){
	return c.dao.Find(id)
}

func (c ServiceSqbUser) FindAll(page components.Page) (*[]*models.SqbUser, error){
	return c.dao.FindAll(page)
}
func (c ServiceSqbUser) FindWithID(ids []uint) (*[]*models.SqbUser, error){
	return c.dao.FindWithID(ids)
}





func (c ServiceSqbUser) CountWhereForAdmin(wh string,bind ...string) (int64,  error) {
     return  c.dao.CountWhereForAdmin(wh, bind...)
}

func (c ServiceSqbUser) FindWhereForAdmin(wh string,page components.Page,bind ...string) (*[]*models.SqbUser, error){
     return  c.dao.FindWhereForAdmin(wh,page, bind...)
}
func (c ServiceSqbUser) Count() (int64, error) {
	return c.dao.Count()
}


func NewServiceSqbUser(conf *config.ServiceConfig) ServiceSqbUser {
	engine := conf.ConnectGroup.MysqlGroup["site_master"].(*xorm.Engine)

    return ServiceSqbUser{conf: conf, dao: daos.NewSqbUserDao(conf,engine)}
}

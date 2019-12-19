package daos

import (
     
        "errors"
     
        "duoduo.com/go-xorm/xorm"
     
        "duoduo.com/micro/common/components"
     
        "duoduo.com/micro/hello/srv/config"
     
        "duoduo.com/micro/hello/srv/models"
     
)



type SqbUserDao struct {
    conf *config.ServiceConfig
    db *xorm.Engine
}
func (c SqbUserDao) Create(model models.SqbUser)   (models.SqbUser, error)  {
    _,err := c.db.Insert(&model)
   	return model, err
}
func (c SqbUserDao) Delete(model models.SqbUser)   (int, error)  {
  	if model.Id  <= 0 {
  		return 0,errors.New("id miss")
  	}
  	affected, err := c.db.Id(model.Id).Delete(&model)
  	if affected > 0 {
  		return int(affected), err
  	}
  	return 0, err
}

func (c SqbUserDao) DeleteById(id uint) (int, error) {
   model, err := c.Find(id)
   	affected, err := c.db.Id(id).Delete(model)
   	if (affected > 0) {
   		return int(affected), nil
   	}
   	return 0, err
}

func (c SqbUserDao) Update(id uint, attr map[string]interface{})   (bool, error)  {
    if id >= 0 {
    		_, err := c.db.Id(id).Table(models.SqbApp{}).Update(attr)
    		if err != nil {
    			return false, err;
    		} else {
    			return true, err
    		}
    }
    return false, errors.New("id miss")
}
func (c SqbUserDao) UpdateByModel(model models.SqbUser)   (bool, error)  {
    if model.Id >= 0 {
    		_, err := c.db.Id(model.Id).Update(model)
    		if err != nil {
    			return false, err;
    		} else {
    			return true, err
    		}
    }
    return false, errors.New("id miss")
}

func (c SqbUserDao) Find(id uint)  ( *models.SqbUser, error) {
    var object models.SqbUser
    model := &object
    _, err := c.db.Id(id).Get(model)
    if err != nil {
        return nil, err
    }
    if model.Id > 0 {
        return model, nil
    }
    return nil, nil
}


func (c SqbUserDao) FindByKey(name string,value string) (*models.SqbUser, error) {
    model := &models.SqbUser{}
   	bool,err := c.db.Where(name+" = ?", value).Limit(1).Get(model)
   	if err == nil {
   		if bool {
   			return model, nil
   		}else{
   			return nil, nil
   		}
   	} else {
   		return model, err
   	}
}



func (c SqbUserDao) FindAll(page components.Page) (*[]*models.SqbUser, error){
    var list []*models.SqbUser
   	limitInfo := page.ToLimitInfo()
   	err := c.db.OrderBy(page.ToOrder()).Limit(limitInfo.Limit, limitInfo.Offset).Find(&list)
   	if len(list) <= 0  {
   		return nil,nil
   	}
   	return &list, err
}
func (c SqbUserDao) FindWithID(ids []uint) (*[]*models.SqbUser, error) {
        var list []*models.SqbUser
     	err := c.db.In("id",ids).Find(&list)
     	if len(list)  <= 0  {
     		return nil,nil
     	}
     	return &list, err
}

func (c SqbUserDao) CountWhereForAdmin(wh string,bind ...string) (int64,  error) {
    model := new(models.SqbUser)
   	return  c.db.Where(wh, bind).Count(model)
}

func (c SqbUserDao) FindWhereForAdmin(wh string,page components.Page,bind ...string)  (*[]*models.SqbUser, error) {
        limitInfo := page.ToLimitInfo()
    	var list []*models.SqbUser
    	err := c.db.Where(wh,bind).OrderBy(page.Order).Limit(limitInfo.Limit,limitInfo.Offset).Find(&list)
    	if len(list)  <= 0  {
    		return nil,nil
    	}
    	return &list, err
}



func (c SqbUserDao) Count() (int64, error) {
    return c.db.Count(&models.SqbUser{})
}

func NewSqbUserDao(conf *config.ServiceConfig, db *xorm.Engine) *SqbUserDao {
    return &SqbUserDao{conf: conf, db: db}
}

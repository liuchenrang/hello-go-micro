package daos

import (
     
        "github.com/pkg/errors"
     
        "duoduo.com/go-xorm/xorm"
     
        "duoduo.com/micro/common/components"
     
        "duoduo.com/micro/hello/srv/config"
     
        "duoduo.com/micro/hello/srv/models"
     
)


var (
    SqbApp SqbAppDao
)
type SqbAppDao struct {
    conf *config.ServiceConfig
    db *xorm.Engine
}
func (c SqbAppDao) Create(model models.SqbApp)   (models.SqbApp, error)  {
    _,err := c.db.Insert(&model)
   	return model, err
}
func (c SqbAppDao) Delete(model models.SqbApp)   (int, error)  {
  	if model.Id  <= 0 {
  		return 0,errors.New("id miss")
  	}
  	affected, err := c.db.Id(model.Id).Delete(&model)
  	if affected > 0 {
  		return int(affected), err
  	}
  	return 0, err
}

func (c SqbAppDao) DeleteById(id uint) (int, error) {
   model, err := c.Find(id)
   	affected, err := c.db.Id(id).Delete(model)
   	if (affected > 0) {
   		return int(affected), nil
   	}
   	return 0, err
}

func (c SqbAppDao) Update(id uint, attr map[string]interface{})   (bool, error)  {
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
func (c SqbAppDao) UpdateByModel(model models.SqbApp)   (bool, error)  {
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

func (c SqbAppDao) Find(id uint)  ( *models.SqbApp, error) {
    var object models.SqbApp
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


func (c SqbAppDao) FindByKey(name string,value string) (*models.SqbApp, error) {
    model := &models.SqbApp{}
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



func (c SqbAppDao) FindAll(page components.Page) (*[]*models.SqbApp, error){
    var list []*models.SqbApp
   	limitInfo := page.ToLimitInfo()
   	err := c.db.OrderBy(page.ToOrder()).Limit(limitInfo.Limit, limitInfo.Offset).Find(&list)
   	if len(list) <= 0  {
   		return nil,nil
   	}
   	return &list, err
}
func (c SqbAppDao) FindWithID(ids []uint) (*[]*models.SqbApp, error) {
        var list []*models.SqbApp
     	err := c.db.In("id",ids).Find(&list)
     	if len(list)  <= 0  {
     		return nil,nil
     	}
     	return &list, err
}

func (c SqbAppDao) CountWhereForAdmin(wh string,bind ...string) (int64,  error) {
    model := new(models.SqbApp)
   	return  c.db.Where(wh, bind).Count(model)
}

func (c SqbAppDao) FindWhereForAdmin(wh string,page components.Page,bind ...string)  (*[]*models.SqbApp, error) {
        limitInfo := page.ToLimitInfo()
    	var list []*models.SqbApp
    	err := c.db.Where(wh,bind).OrderBy(page.Order).Limit(limitInfo.Limit,limitInfo.Offset).Find(&list)
    	if len(list)  <= 0  {
    		return nil,nil
    	}
    	return &list, err
}



func (c SqbAppDao) Count() (int64, error) {
    return c.db.Count(&models.SqbApp{})
}

func NewSqbAppDao(conf *config.ServiceConfig, db *xorm.Engine) *SqbAppDao {
    return &SqbAppDao{conf: conf, db: db}
}

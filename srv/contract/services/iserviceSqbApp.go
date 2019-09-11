package contractService

import (
	"xiaoshijie.com/micro/common/components"
	
	"xiaoshijie.com/micro/hello/srv/models"
)

type IServiceSqbApp interface {
	Add(model models.SqbApp) (models.SqbApp, error)
	Delete(model models.SqbApp) (int, error)
	DeleteById(id uint) (int, error)
	
	Update(id uint, attr map[string]interface{}) (bool, error)
	UpdateByModel(model models.SqbApp) (bool, error)
	
	FindCode(code string) (*models.SqbApp, error)
	Find(id uint) (*models.SqbApp, error)
	
	FindAll(page components.Page) (*[]*models.SqbApp, error)
	FindWithID(ids []uint) (*[]*models.SqbApp, error)
	
	CountWhereForAdmin(wh string, bind ...string) (int64, error)
	
	FindWhereForAdmin(wh string, page components.Page, bind ...string) (*[]*models.SqbApp, error)
	
	Count() (int64, error)
}

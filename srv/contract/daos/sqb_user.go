package contractDao

import (
    "xiaoshijie.com/micro/common/components"
    
    "xiaoshijie.com/micro/hello/srv/models"
)
type ISqbUserDao interface {
   Create(model models.SqbUser)  (models.SqbUser, error)
   Delete(model models.SqbUser)  (int, error)
   DeleteById(id uint)  (int, error)

   Update(id uint, attr map[string]interface{})  (bool, error)
   UpdateByModel(model models.SqbUser)  (bool, error)


   Find(id uint)  ( *models.SqbUser, error)

   FindByKey(name string,value string) (*models.SqbUser, error)

   FindAll(page components.Page) (*[]*models.SqbUser, error)
   FindWithID(ids []uint) (*[]*models.SqbUser, error)

   CountWhereForAdmin(wh string,bind ...string) (int64,  error)

   FindWhereForAdmin(wh string,page components.Page,bind ...string) (*[]*models.SqbUser, error)

   Count() (int64, error)
}


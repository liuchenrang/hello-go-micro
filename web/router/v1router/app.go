package v1router

import (
	"github.com/gin-gonic/gin"
	"xiaoshijie.com/micro/tbservice/tb-service-web/controller/v1controller"
)

func InitV1Router(r gin.IRoutes)  {
	r.GET("1/app/guideInfo", v1controller.App.GuideInfo)
	r.POST("1/collect/push", v1controller.Collect.Push)
}

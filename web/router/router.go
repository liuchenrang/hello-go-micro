package router

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"xiaoshijie.com/micro/common/http/middleware"
	"xiaoshijie.com/micro/tbservice/tb-service-web/router/v1router"
	"xiaoshijie.com/micro/tbservice/tb-service-web/router/v3router"
	"xiaoshijie.com/micro/tbservice/tb-service-web/router/v4router"
)
var (
	//r *gin.Engine
	r = InitRouter()
)

func InitRouter() *gin.Engine  {
	engine := gin.New()
	engine.Use(gin.Recovery())
	logger := &lumberjack.Logger{
		Filename:   "/tmp/micro-tb-service-web.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     3,    //days
		Compress:   true, // disabled by default
	}
	
	engine.Use(gin.LoggerWithWriter(logger))
	return engine
}
func init()  {

}
func InitApiRouter(r  *gin.Engine )  {
	apiGroup := r.Group("api/").Use(middleware.ApiSignChecker()).Use(middleware.ApiTokenChecker())
	v1router.InitV1Router(apiGroup)
	v3router.InitV3Router(apiGroup)
	v4router.InitV4Router(apiGroup)
	
}
func GetRouter() *gin.Engine  {
	
	InitApiRouter(r)
	return r
}
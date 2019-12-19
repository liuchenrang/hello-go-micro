module xiaoshijie.com/micro/tbservice/tb-service-web

go 1.12

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.3.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0-20170531160350-a96e63847dc3
	xiaoshijie.com/micro/common v0.0.0-00010101000000-000000000000
	xiaoshijie.com/micro/tbservice/tb-service-srv v0.0.0-00010101000000-000000000000
)

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.3.0

replace xiaoshijie.com/micro/tbservice/tb-service-srv => ../tb-service-srv

replace xiaoshijie.com/micro/common => ../../common

replace xiaoshijie.com/go-xorm/xorm => ../../customLib/go-xorm/xorm

replace xiaoshijie.com/go-xorm/core => ../../customLib/go-xorm/core

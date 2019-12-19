module xiaoshijie.com/micro/tbservice/tb-service-cmd

go 1.12

require (
	github.com/gogo/protobuf v1.2.1
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.3.0
	github.com/micro/micro v1.14.0
	xiaoshijie.com/micro/common v0.0.0-00010101000000-000000000000
	xiaoshijie.com/micro/tbservice/tb-service-srv v0.0.0-00010101000000-000000000000
)

replace xiaoshijie.com/micro/tbservice/tb-service-srv => ../tb-service-srv

replace xiaoshijie.com/micro/common => ../../common

replace xiaoshijie.com/go-xorm/xorm => ../../customLib/go-xorm/xorm

replace xiaoshijie.com/go-xorm/core => ../../customLib/go-xorm/core

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.3.0

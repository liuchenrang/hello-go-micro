module xiaoshijie.com/micro/hello/srv

go 1.12

require (
	github.com/go-redis/redis v6.15.5+incompatible // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.9.1
	github.com/micro/go-plugins v1.2.0
	github.com/pkg/errors v0.8.1
	xiaoshijie.com/go-xorm/core v0.0.0-00010101000000-000000000000 // indirect
	xiaoshijie.com/go-xorm/xorm v0.0.0-00010101000000-000000000000 // indirect
	xiaoshijie.com/micro/common v0.0.0-00010101000000-000000000000

)

replace xiaoshijie.com/go-xorm/xorm => ../customLib/go-xorm/xorm

replace xiaoshijie.com/go-xorm/core => ../customLib/go-xorm/core

replace xiaoshijie.com/micro/common => ../../common

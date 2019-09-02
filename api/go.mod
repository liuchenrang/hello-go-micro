module xiaoshijie.com/micro/hello/api

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.9.1
	github.com/micro/go-plugins v1.2.0
	xiaoshijie.com/micro/hello/srv v0.0.0-00010101000000-000000000000
)

replace xiaoshijie.com/micro/hello/srv => ../srv

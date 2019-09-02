###编译事项
    编译遇到问题删除包冲突
###启动api代理
   ./micro  --registry=etcdv3  --registry_address=http://0.0.0.0:2379  api  --address 0.0.0.0:8181  --handler api


###启动api服务
   GO111MODULE=on MICRO_CLIENT=grpc  MICRO_SERVER=api go run main.go plugins.go --registry=etcdv3

###启动srv服务
   GO111MODULE=on go run main.go plugins.go --registry=etcdv3  --registry_address=http://0.0.0.0:2379

###代码生成
   micro  new xiaoshijie.com/micro/hello/cli --type=cli
###go-etcd-配置中心
   配置中心的client的代码采用的是etcdv3接口。
###go-etcd-服务注册
   注册中心采用的是 etcdv3接口
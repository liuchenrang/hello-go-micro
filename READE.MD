删除vendor下 冲突目录

启动api代理

./micro  --registry=etcdv3  --registry_address=http://0.0.0.0:2379  api  --address 0.0.0.0:8181  --handler api


启动api服务

  GO111MODULE=on MICRO_CLIENT=grpc  MICRO_SERVER=api go run main.go plugins.go --registry=etcdv3

启动srv服务

 GO111MODULE=on go run main.go plugins.go --registry=etcdv3  --registry_address=http://0.0.0.0:2379

代码生成

micro  new xiaoshijie.com/micro/hello/cli --type=cli
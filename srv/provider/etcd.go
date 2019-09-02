package provider

import (
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-plugins/config/source/etcd"
)
var (
	EtcdSource source.Source
)
func init()  {
	println("ectd source init")
	EtcdSource = etcd.NewSource(
		// optionally specify etcd address; default to localhost:8500
		etcd.WithAddress("192.168.1.122:12379"),
		// optionally specify prefix; defaults to /micro/config
		//etcd.WithPrefix("/config-center/development/project/prefix"),
		etcd.WithPrefix("/micro/config"),
		// optionally strip the provided prefix from the keys, defaults to false
		etcd.StripPrefix(true),
	)
}

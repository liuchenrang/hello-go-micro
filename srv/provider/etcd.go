package provider

import (
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-plugins/config/source/etcd"
)
var (
	EtcdSource source.Source
)
func init()  {
	EtcdSource = etcd.NewSource(
		// optionally specify etcd address; default to localhost:8500
		etcd.WithAddress("127.0.0.1:2379"),
		// optionally specify prefix; defaults to /micro/config
		etcd.WithPrefix("/config-center/development/project/prefix"),
		// optionally strip the provided prefix from the keys, defaults to false
		etcd.StripPrefix(true),
	)
}

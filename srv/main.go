package main

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"jun.srv.id/proto"
	"jun.srv.id/srv/handler"
)

var (
	regAddr = "106.13.214.251:2379"
	name    = "jun.srv.id"
	version = "latest"
	service micro.Service
	err     error
)

func main() {
	service = micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Registry(etcd.NewRegistry(registry.Addrs(regAddr))),
	)

	service.Init()
	_ = proto.RegisterIdServiceHandler(service.Server(), handler.NewIdHandler())
	if err = service.Run(); err != nil {
		panic(fmt.Errorf("开启Id微服务失败: err=%v", err))
	}
}

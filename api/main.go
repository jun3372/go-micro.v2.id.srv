package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"github.com/spf13/cast"

	"jun.srv.id/proto"
)

var (
	regAddr = "106.13.214.251:2379"
	name    = "jun.srv.id.client"
	addr    = ":8001"
	version = "latest"
	srvName = "jun.srv.id"

	service web.Service
	req     proto.IdRequest
	res     *proto.IdResponse
	err     error
)

func main() {
	service = web.NewService(
		web.Name(name),
		web.Version(version),
		web.Address(addr),
		web.Registry(etcd.NewRegistry(registry.Addrs(regAddr))),
	)

	engine := gin.Default()
	engine.Any("/", func(ctx *gin.Context) {
		if err = ctx.Bind(&req); err != nil {
			log.Debug("获取参数失败: err=", err)
			ctx.JSON(200, gin.H{"err": 1, "msg": "获取参数失败"})
			return
		}

		if req.Node < 1 {
			req.Node = cast.ToInt64(ctx.PostForm("node"))
		}

		idService := proto.NewIdService(srvName, client.DefaultClient)
		if res, err = idService.GetId(context.TODO(), &req); err != nil {
			log.Debug("获取Id失败: err=", err)
			ctx.JSON(200, gin.H{"err": 1, "msg": fmt.Sprintf("获取Id失败, err=%s", err)})
			return
		}
		log.Info("res=", res)
		ctx.JSON(200, gin.H{"err": 0, "msg": "获取Id成功", "data": res})
	})

	service.Handle("/", engine)
	service.Init()
	if err = service.Run(); err != nil {
		panic(fmt.Errorf("启动web服务失败: err=%v", err))
	}
}

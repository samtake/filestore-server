package main

import (
	"fmt"
	"time"

	micro "github.com/micro/go-micro"

	cfg "filestore-server/service/download/config"
	dlProto "filestore-server/service/download/proto"
	"filestore-server/service/download/route"
	dlRpc "filestore-server/service/download/rpc"
)

func startRpcService() {
	service := micro.NewService(
		micro.Name("go.micro.service.download"), // 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	service.Init()

	dlProto.RegisterDownloadServiceHandler(service.Server(), new(dlRpc.Download))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startApiService() {
	router := route.Router()
	router.Run(cfg.DownloadServiceHost)
}

func main() {
	// api 服务
	go startApiService()

	// rpc 服务
	startRpcService()
}

package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"

	dbProxy "filestore-server/service/dbproxy/proto"
	dbRpc "filestore-server/service/dbproxy/rpc"
)

func startRpcService() {
	service := micro.NewService(
		micro.Name("go.micro.service.dbproxy"), // 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),      // 声明超时时间, 避免consul不主动删掉已失去心跳的服务节点
		micro.RegisterInterval(time.Second*5),
	)
	service.Init()

	dbProxy.RegisterDBProxyServiceHandler(service.Server(), new(dbRpc.DBProxy))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	startRpcService()
	// res, err := mapper.FuncCall("/user/UserExist", []interface{}{"haha"}...)
	// log.Printf("error: %+v\n", err)
	// log.Printf("result: %+v\n", res[0].Interface())

	// res, err = mapper.FuncCall("/user/UserExist", []interface{}{"admin"}...)
	// log.Printf("error: %+v\n", err)
	// log.Printf("result: %+v\n", res[0].Interface())
}

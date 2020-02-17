package main

import (
	"filestore-server/config"
	"filestore-server/route"
	"fmt"
)

func main() {
	// gin framework
	router := route.Router()

	// 启动服务并监听端口
	err := router.Run(config.UploadServiceHost)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s\n", err.Error())
	}
}

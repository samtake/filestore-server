package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"filestore-server/config"
	"filestore-server/mq"
	dbcli "filestore-server/service/dbproxy/client"
	"filestore-server/store/oss"

	micro "github.com/micro/go-micro"
)

// ProcessTransfer : 处理文件转移
func ProcessTransfer(msg []byte) bool {
	log.Println(string(msg))

	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	fin, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	err = oss.Bucket().PutObject(
		pubData.DestLocation,
		bufio.NewReader(fin))
	if err != nil {
		log.Println(err.Error())
		return false
	}

	resp, err := dbcli.UpdateFileLocation(
		pubData.FileHash,
		pubData.DestLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if !resp.Suc {
		log.Println("更新数据库异常，请检查:" + pubData.FileHash)
		return false
	}
	return true
}

func startRpcService() {
	service := micro.NewService(
		micro.Name("go.micro.service.transfer"), // 服务名称
		micro.RegisterTTL(time.Second*10),       // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5),   // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
	)
	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startTranserService() {
	if !config.AsyncTransferEnable {
		log.Println("异步转移文件功能目前被禁用，请检查相关配置")
		return
	}
	log.Println("文件转移服务启动中，开始监听转移任务队列...")
	mq.StartConsume(
		config.TransOSSQueueName,
		"transfer_oss",
		ProcessTransfer)
}

func main() {
	// 文件转移服务
	go startTranserService()

	// rpc 服务
	startRpcService()
}

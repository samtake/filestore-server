package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	//N设定路由规则
	http.HandleFunc("/file/upload", handler.UploadHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed to start server,err:%s ", err.Error())
	}
}

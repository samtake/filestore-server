package handler

import (
	"encoding/json"
	"filestore-server/meta"
	"filestore-server/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//文件上传
func UploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			//返回上传的html页面
			io.WriteString(w, "internet server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收文件流及存储到本地
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data err : %s\n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp" + head.Filename,
			UploadAt: time.Now().Format("2001-11-12 12:12:12"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("failed to create file,err: %s", err.Error())
			return
		}
		// defer newFile.Close()
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("failed to save file,err: %s", err.Error())
			return
		}

		http.Redirect(w, r, "file/upload/suc", http.StatusFound)
	}
}

//上传成功
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Success")
}

//获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

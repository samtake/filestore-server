# distributed-fileserver

基于golang实现的一种分布式云存储服务


## 静态资源打包

```bash
# 下载库
go get -v github.com/jteeuwen/go-bindata/...
go get -v github.com/moxiaomomo/go-bindata-assetfs/...

# 将$GOPATH/bin关联到$PATH中, 可修改~/.bashrc文件(go-bindata-assetfs命令安装在$GOPATH/bin下)
export PATH=$PATH:$GOPATH/bin

# cd $GOPATH/<你的工程目录>
cd $GOPATH/filestore-server

# 将静态文件打包到一个目标文件里
mkdir assets -p && go-bindata-assetfs -pkg assets -o ./assets/asset.go static/...

# 修改静态文件的处理逻辑，详细可参考./service/upload/main.go
```
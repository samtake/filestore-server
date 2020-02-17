# distributed-fileserver
基于golang实现的一种分布式云存储服务(当前分支是用Gin框架改造的一个版本)

## 关于安装gin包

如果有梯子，直接可以下

```shell
go get github.com/gin-gonic/gin
```

失败的话可尝试以下这些步骤下载:

```shell
mkdir $GOPATH/src/golang.org/x -p
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/sys

mkdir $GOPATH/src/gopkg.in/go-playground -p
cd $GOPATH/src/gopkg.in/go-playground
git clone -b v8 https://github.com/go-playground/validator.git
mv validator validator.v8

cd $GOPATH/src/gopkg.in/
git clone -b v2 https://github.com/go-yaml/yaml.git
mv yaml yaml.v2

cd $GOPATH
go get github.com/gin-gonic/gin
```

## 进度说明：
* [x] 简单的文件上传服务
* [x] mysql存储文件元数据
* [x] 账号系统, 注册/登录/查询用户或文件数据
* [x] 基于帐号的文件操作接口
* [x] 文件秒传功能
* [x] 文件分块上传/断点续传功能
* [x] 搭建及使用Ceph对象存储集群
* [x] 使用阿里云OSS对象存储服务
* [ ] 使用RabbitMQ实现异步任务队列
* [ ] 微服务化(API网关, 服务注册, RPC通讯)
* [ ] CI/CD(持续集成)

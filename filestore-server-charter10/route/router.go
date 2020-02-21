package route

import (
	"filestore-server/handler"

	"github.com/gin-gonic/gin"
)

// Router : 路由表配置
func Router() *gin.Engine {
	// gin framework, 包括Logger, Recovery
	router := gin.Default()

	// 处理静态资源
	router.Static("/static/", "./static")

	// 不需要经过验证就能访问的接口
	router.GET("/user/signup", handler.SignupHandler)
	router.POST("/user/signup", handler.DoSignupHandler)

	router.GET("/user/signin", handler.SignInHandler)
	router.POST("/user/signin", handler.DoSignInHandler)

	// 加入中间件，用于校验token的拦截器
	router.Use(handler.HTTPInterceptor())

	// Use之后的所有handler都会经过拦截器进行token校验

	// 文件存取接口
	router.GET("/file/upload", handler.UploadHandler)
	router.POST("/file/upload", handler.DoUploadHandler)
	router.GET("/file/upload/suc", handler.UploadSucHandler)
	router.POST("/file/meta", handler.GetFileMetaHandler)
	router.POST("/file/query", handler.FileQueryHandler)
	router.GET("/file/download", handler.DownloadHandler)
	router.POST("/file/update", handler.FileMetaUpdateHandler)
	router.POST("/file/delete", handler.FileDeleteHandler)
	router.POST("/file/downloadurl", handler.DownloadURLHandler)

	// 秒传接口
	router.POST("/file/fastupload", handler.TryFastUploadHandler)

	// 分块上传接口
	router.POST("/file/mpupload/init", handler.InitialMultipartUploadHandler)
	router.POST("/file/mpupload/uppart", handler.UploadPartHandler)
	router.POST("/file/mpupload/complete", handler.CompleteUploadHandler)

	// 用户相关接口
	router.POST("/user/info", handler.UserInfoHandler)

	return router
}

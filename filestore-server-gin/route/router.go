package route

import (
	hdl "filestore-server/handler"

	"github.com/gin-gonic/gin"
)

// Router ：路由规则定义
func Router() *gin.Engine {
	// gin framework
	router := gin.Default()

	// 静态资源处理
	router.Static("/static/", "./static")

	// 不需验证的接口
	router.GET("/user/signup", hdl.SignupHandler)
	router.GET("/user/signin", hdl.SigninHandler)
	router.POST("/user/signup", hdl.DoSignupHandler)
	router.POST("/user/signin", hdl.DoSignInHandler)
	router.GET("/user/exists", hdl.UserExistsHandler)

	// 加入auth认证中间件
	router.Use(hdl.Authorize())

	// 文件存取接口
	router.GET("/file/upload", hdl.UploadHandler)
	router.POST("/file/upload", hdl.DoUploadHandler)
	router.GET("/file/upload/suc", hdl.UploadSucHandler)
	router.GET("/file/meta", hdl.GetFileMetaHandler)
	router.POST("/file/query", hdl.FileQueryHandler)
	router.GET("/file/download", hdl.DownloadHandler)
	router.POST("/file/download", hdl.DownloadHandler)
	router.POST("/file/update", hdl.FileMetaUpdateHandler)
	router.POST("/file/delete", hdl.FileDeleteHandler)
	router.POST("/file/downloadurl",
		hdl.DownloadURLHandler)

	// 秒传接口
	router.POST("/file/fastupload",
		hdl.TryFastUploadHandler)

	// 分块上传接口
	router.POST("/file/mpupload/init",
		hdl.InitialMultipartUploadHandler)
	router.POST("/file/mpupload/uppart",
		hdl.UploadPartHandler)
	router.POST("/file/mpupload/complete",
		hdl.CompleteUploadHandler)

	// 用户相关接口
	router.POST("/user/info", hdl.UserInfoHandler)

	return router
}

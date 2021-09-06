package router

import (
	v1 "Hyaenidae/api/v1"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup) {
	FileRouter := Router.Group("fileUploadAndDownload")
	{
		FileRouter.POST("uploadFile", v1.UploadFile)       // 保存file
		FileRouter.POST("deleteFile", v1.DeleteFile)       // 保存file

	}
}
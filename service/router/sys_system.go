package router

import (
	v1 "Hyaenidae/api/v1"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system")
	{
		SystemRouter.POST("getServerInfo", v1.GetServerInfo) // 获取服务器信息
	}
}

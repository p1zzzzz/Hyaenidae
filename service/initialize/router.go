package initialize

import (
	"Hyaenidae/global"
	"Hyaenidae/middleware"
	"Hyaenidae/router"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 方便统一添加路由组前缀 多服务器上线使用
	Router.Use(ginzap.Ginzap(global.Hyaenidae_LOG, time.RFC3339, true))
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		router.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		//router.InitApiRouter(PrivateGroup)                   // 注册功能api路由
		router.InitJwtRouter(PrivateGroup)       // jwt相关路由
		router.InitUserRouter(PrivateGroup)      // 注册用户路由
		router.InitAuthorityRouter(PrivateGroup) //注册角色路由
		router.InitMenuRouter(PrivateGroup)      // 注册menu路由
		router.InitSystemRouter(PrivateGroup)    //注册系统路由
		router.InitMarkdownRouter(PrivateGroup)    //注册markdown路由
		router.InitFileRouter(PrivateGroup)    //注册文件路由
		router.InitExaRouter(PrivateGroup)


	}
	global.Hyaenidae_LOG.Info("router register success")
	return Router
}

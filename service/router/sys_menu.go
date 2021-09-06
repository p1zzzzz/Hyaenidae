package router

import (
	v1 "Hyaenidae/api/v1"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu")
	{
		MenuRouter.POST("getMenu", v1.GetMenu)                   // 获取菜单树
		MenuRouter.POST("getBaseMenuTree", v1.GetBaseMenuTree)   // 获取用户动态路由
		MenuRouter.POST("getMenuAuthority", v1.GetMenuAuthority) //获取指定角色menu
		MenuRouter.POST("addMenuAuthority",v1.AddMenuAuthority)//增加menu和角色关联关系
	}
	return MenuRouter
}

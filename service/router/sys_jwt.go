package router

import (
	v1 "Hyaenidae/api/v1"
	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	JwtRouter := Router.Group("jwt")
	{
		JwtRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}

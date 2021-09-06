package router

import (
	v1 "Hyaenidae/api/v1"
	"github.com/gin-gonic/gin"
)

func InitExaRouter(Router *gin.RouterGroup) {
	ExaRouter := Router.Group("exa")
	{
		ExaRouter.POST("whatAv", v1.WhatAv)   // 查询杀毒进程
	}
}

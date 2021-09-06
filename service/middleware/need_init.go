package middleware

import (
	"Hyaenidae/global"
	"Hyaenidae/model/response"
	"github.com/gin-gonic/gin"
)

func NeedInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.Hyaenidae_DB == nil {
			response.OkWithDetailed(gin.H{
				"needInit": true,
			}, "前往初始化数据库", c)
			c.Abort()
		} else {
			c.Next()
		}
		// 处理请求
	}
}

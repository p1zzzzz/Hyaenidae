package router

import (
	v1 "Hyaenidae/api/v1"
	"github.com/gin-gonic/gin"
)

func InitMarkdownRouter(Router *gin.RouterGroup) {
	MdRouter := Router.Group("markdown")
	{
		MdRouter.POST("saveMd", v1.SaveMd)       // 保存markdown
		MdRouter.POST("deleteMd", v1.DeleteMd) // 删除markdown
		MdRouter.POST("getMdTree", v1.GetMdTree) // 获取markdown树
		MdRouter.POST("getMd", v1.GetMd) // 获取markdown树

	}
}

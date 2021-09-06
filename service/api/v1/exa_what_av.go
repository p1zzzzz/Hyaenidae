package v1

import (
	"Hyaenidae/model/request"
	"Hyaenidae/model/response"
	"Hyaenidae/service"
	"Hyaenidae/utils"
	"github.com/gin-gonic/gin"
)

// @Tags Exa_WhatAv
// @Summary 查询杀毒软件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TaskList true "tasklist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exa/whatAv [post]
func WhatAv(c *gin.Context) {
	var TaskList request.TaskList
	_ = c.ShouldBindJSON(&TaskList)
	if err := utils.Verify(TaskList, utils.TaskListverify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list := service.WhatAv(TaskList.SearchKey)
	response.OkWithDetailed(response.PageResult{List:list}, "查询成功", c)
}



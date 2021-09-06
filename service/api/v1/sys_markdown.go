package v1

import (
	"Hyaenidae/global"
	"Hyaenidae/model"
	"Hyaenidae/model/request"
	"Hyaenidae/model/response"
	"Hyaenidae/service"
	"Hyaenidae/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags SysMarkdown
// @Summary 用户保存markdown
// @Produce  application/json
// @Param data body model.Markdown true "编辑者, 标题, 文章内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"保存成功"}"
// @Router /markdown/saveMd [post]
func SaveMd(c *gin.Context) {
	var r request.SaveMd
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.MarkdownVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mark := &model.Markdown{UUID: r.UUID,Editor: r.Editor, Title: r.Title, MkValue: r.MkValue,ParentId: r.ParentId}
	err, markReturn := service.SaveMd(*mark)
	if err != nil {
		global.Hyaenidae_LOG.Error("保存失败!", zap.Any("err", err))
		response.FailWithDetailed(response.SysMarkdownResponse{Markdown: markReturn}, "保存失败", c)
	} else {
		response.OkWithDetailed(response.SysMarkdownResponse{Markdown: markReturn}, "保存成功", c)
	}
}

// @Tags SysMarkdown
// @Summary 删除markdown
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "MarkdownID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /markdown/deleteMd [delete]
func DeleteMd(c *gin.Context) {
	var reqUuid request.GetByUuid
	_ = c.ShouldBindJSON(&reqUuid)
	if err := utils.Verify(reqUuid, utils.Uuidverify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteMd(reqUuid.UUID); err != nil {
		global.Hyaenidae_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysMarkdown
// @Summary 查询markdown
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "MarkdownID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /markdown/getMd [delete]
func GetMd(c *gin.Context) {
	var reqUuid request.GetByUuid
	_ = c.ShouldBindJSON(&reqUuid)
	if err := utils.Verify(reqUuid, utils.Uuidverify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err,markReturn := service.GetMdByUuid(reqUuid.UUID); err != nil {
		global.Hyaenidae_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.SysMarkdownResponse{Markdown: *markReturn}, "查询成功", c)
	}
}


// @Tags SysMarkdown
// @Summary 获取md树
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /markdown/getMdTree [post]
func GetMdTree(c *gin.Context) {
	if err, mds := service.GetMdTree(); err != nil {
		global.Hyaenidae_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysMarkdownsResponse{Markdowns: mds}, "获取成功", c)
	}
}

import service from '@/utils/request'
import fileService from '@/utils/fileRequest'


// @Tags FileUploadAndDownload
// @Summary 上传文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dbModel.FileUploadAndDownload true 
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/uploadFile [post]
export const uploadFile = (data) => {
    return fileService({
      url: '/fileUploadAndDownload/uploadFile',
      method: 'post',
      data
    })
  }



// @Tags FileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dbModel.FileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
export const deleteFile = (data) => {
    return service({
      url: '/fileUploadAndDownload/deleteFile',
      method: 'post',
      data
    })
  }
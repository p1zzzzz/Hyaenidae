import service from '@/utils/request'

// @Summary 查询杀毒软件
// @Produce  application/json
// @Param data body {taskList:"string"}
// @Router /exa/whatav [post]
export const whatav = (data) => {
    return service({
      url: '/exa/whatAv',
      method: 'post',
      data: data
    })
  }
import service from '@/utils/request'

// @Summary 用户保存markdown
// @Produce  application/json
// @Param data body {editor:"string",title:"string",mkvalue:"string"}
// @Router /markdown/saveMd [post]
export const savemd = (data) => {
    return service({
      url: '/markdown/saveMd',
      method: 'post',
      data: data
    })
  }


// @Summary  获取markdown树
// @Produce  application/json
// @Param 可以什么都不填 调一下即可
// @Router /menu/getMdTree [post]
export const getmdtree = () => {
  return service({
    url: '/markdown/getMdTree',
    method: 'post'
  })
}


// @Summary  获取markdown
// @Produce  application/json
// @Param uuid string
// @Router /menu/getMd [post]
export const getmd = (data) => {
  return service({
    url: '/markdown/getMd',
    method: 'post',
    data :data
  })
}

// @Summary  删除markdown
// @Produce  application/json
// @Param uuid string
// @Router /menu/deleteMd [post]
export const deletemd = (data) => {
  return service({
    url: '/markdown/deleteMd',
    method: 'post',
    data :data
  })
}
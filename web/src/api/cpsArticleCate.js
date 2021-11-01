import service from '@/utils/request'

// @Tags CpsArticleCate
// @Summary 创建CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CpsArticleCate true "创建CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpsArticleCate/createCpsArticleCate [post]
export const createCpsArticleCate = (data) => {
  return service({
    url: '/cpsArticleCate/createCpsArticleCate',
    method: 'post',
    data
  })
}

// @Tags CpsArticleCate
// @Summary 删除CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CpsArticleCate true "删除CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpsArticleCate/deleteCpsArticleCate [delete]
export const deleteCpsArticleCate = (data) => {
  return service({
    url: '/cpsArticleCate/deleteCpsArticleCate',
    method: 'delete',
    data
  })
}

// @Tags CpsArticleCate
// @Summary 删除CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpsArticleCate/deleteCpsArticleCate [delete]
export const deleteCpsArticleCateByIds = (data) => {
  return service({
    url: '/cpsArticleCate/deleteCpsArticleCateByIds',
    method: 'delete',
    data
  })
}

// @Tags CpsArticleCate
// @Summary 更新CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CpsArticleCate true "更新CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cpsArticleCate/updateCpsArticleCate [put]
export const updateCpsArticleCate = (data) => {
  return service({
    url: '/cpsArticleCate/updateCpsArticleCate',
    method: 'put',
    data
  })
}

// @Tags CpsArticleCate
// @Summary 用id查询CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CpsArticleCate true "用id查询CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cpsArticleCate/findCpsArticleCate [get]
export const findCpsArticleCate = (params) => {
  return service({
    url: '/cpsArticleCate/findCpsArticleCate',
    method: 'get',
    params
  })
}

// @Tags CpsArticleCate
// @Summary 分页获取CpsArticleCate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CpsArticleCate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpsArticleCate/getCpsArticleCateList [get]
export const getCpsArticleCateList = (params) => {
  return service({
    url: '/cpsArticleCate/getCpsArticleCateList',
    method: 'get',
    params
  })
}

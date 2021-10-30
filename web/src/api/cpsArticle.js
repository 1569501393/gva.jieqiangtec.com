import service from '@/utils/request'

// @Tags CpsArticle
// @Summary 创建CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CpsArticle true "创建CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpsArticle/createCpsArticle [post]
export const createCpsArticle = (data) => {
  return service({
    url: '/cpsArticle/createCpsArticle',
    method: 'post',
    data
  })
}

// @Tags CpsArticle
// @Summary 删除CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CpsArticle true "删除CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpsArticle/deleteCpsArticle [delete]
export const deleteCpsArticle = (data) => {
  return service({
    url: '/cpsArticle/deleteCpsArticle',
    method: 'delete',
    data
  })
}

// @Tags CpsArticle
// @Summary 删除CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpsArticle/deleteCpsArticle [delete]
export const deleteCpsArticleByIds = (data) => {
  return service({
    url: '/cpsArticle/deleteCpsArticleByIds',
    method: 'delete',
    data
  })
}

// @Tags CpsArticle
// @Summary 更新CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CpsArticle true "更新CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cpsArticle/updateCpsArticle [put]
export const updateCpsArticle = (data) => {
  return service({
    url: '/cpsArticle/updateCpsArticle',
    method: 'put',
    data
  })
}

// @Tags CpsArticle
// @Summary 用id查询CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CpsArticle true "用id查询CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cpsArticle/findCpsArticle [get]
export const findCpsArticle = (params) => {
  return service({
    url: '/cpsArticle/findCpsArticle',
    method: 'get',
    params
  })
}

// @Tags CpsArticle
// @Summary 分页获取CpsArticle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CpsArticle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpsArticle/getCpsArticleList [get]
export const getCpsArticleList = (params) => {
  return service({
    url: '/cpsArticle/getCpsArticleList',
    method: 'get',
    params
  })
}

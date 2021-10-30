package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CpsArticleApi struct {
}

var cpsArticleService = service.ServiceGroupApp.AutoCodeServiceGroup.CpsArticleService


// CreateCpsArticle 创建CpsArticle
// @Tags CpsArticle
// @Summary 创建CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CpsArticle true "创建CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpsArticle/createCpsArticle [post]
func (cpsArticleApi *CpsArticleApi) CreateCpsArticle(c *gin.Context) {
	var cpsArticle autocode.CpsArticle
	_ = c.ShouldBindJSON(&cpsArticle)
	if err := cpsArticleService.CreateCpsArticle(cpsArticle); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCpsArticle 删除CpsArticle
// @Tags CpsArticle
// @Summary 删除CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CpsArticle true "删除CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpsArticle/deleteCpsArticle [delete]
func (cpsArticleApi *CpsArticleApi) DeleteCpsArticle(c *gin.Context) {
	var cpsArticle autocode.CpsArticle
	_ = c.ShouldBindJSON(&cpsArticle)
	if err := cpsArticleService.DeleteCpsArticle(cpsArticle); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCpsArticleByIds 批量删除CpsArticle
// @Tags CpsArticle
// @Summary 批量删除CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cpsArticle/deleteCpsArticleByIds [delete]
func (cpsArticleApi *CpsArticleApi) DeleteCpsArticleByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := cpsArticleService.DeleteCpsArticleByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCpsArticle 更新CpsArticle
// @Tags CpsArticle
// @Summary 更新CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CpsArticle true "更新CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cpsArticle/updateCpsArticle [put]
func (cpsArticleApi *CpsArticleApi) UpdateCpsArticle(c *gin.Context) {
	var cpsArticle autocode.CpsArticle
	_ = c.ShouldBindJSON(&cpsArticle)
	if err := cpsArticleService.UpdateCpsArticle(cpsArticle); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCpsArticle 用id查询CpsArticle
// @Tags CpsArticle
// @Summary 用id查询CpsArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CpsArticle true "用id查询CpsArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cpsArticle/findCpsArticle [get]
func (cpsArticleApi *CpsArticleApi) FindCpsArticle(c *gin.Context) {
	var cpsArticle autocode.CpsArticle
	_ = c.ShouldBindQuery(&cpsArticle)
	if err, recpsArticle := cpsArticleService.GetCpsArticle(cpsArticle.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recpsArticle": recpsArticle}, c)
	}
}

// GetCpsArticleList 分页获取CpsArticle列表
// @Tags CpsArticle
// @Summary 分页获取CpsArticle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CpsArticleSearch true "分页获取CpsArticle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpsArticle/getCpsArticleList [get]
func (cpsArticleApi *CpsArticleApi) GetCpsArticleList(c *gin.Context) {
	var pageInfo autocodeReq.CpsArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cpsArticleService.GetCpsArticleInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CpsArticleCateApi struct {
}

var cpsArticleCateService = service.ServiceGroupApp.AutoCodeServiceGroup.CpsArticleCateService


// CreateCpsArticleCate 创建CpsArticleCate
// @Tags CpsArticleCate
// @Summary 创建CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CpsArticleCate true "创建CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpsArticleCate/createCpsArticleCate [post]
func (cpsArticleCateApi *CpsArticleCateApi) CreateCpsArticleCate(c *gin.Context) {
	var cpsArticleCate autocode.CpsArticleCate
	_ = c.ShouldBindJSON(&cpsArticleCate)
	if err := cpsArticleCateService.CreateCpsArticleCate(cpsArticleCate); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCpsArticleCate 删除CpsArticleCate
// @Tags CpsArticleCate
// @Summary 删除CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CpsArticleCate true "删除CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpsArticleCate/deleteCpsArticleCate [delete]
func (cpsArticleCateApi *CpsArticleCateApi) DeleteCpsArticleCate(c *gin.Context) {
	var cpsArticleCate autocode.CpsArticleCate
	_ = c.ShouldBindJSON(&cpsArticleCate)
	if err := cpsArticleCateService.DeleteCpsArticleCate(cpsArticleCate); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCpsArticleCateByIds 批量删除CpsArticleCate
// @Tags CpsArticleCate
// @Summary 批量删除CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cpsArticleCate/deleteCpsArticleCateByIds [delete]
func (cpsArticleCateApi *CpsArticleCateApi) DeleteCpsArticleCateByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := cpsArticleCateService.DeleteCpsArticleCateByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCpsArticleCate 更新CpsArticleCate
// @Tags CpsArticleCate
// @Summary 更新CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CpsArticleCate true "更新CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cpsArticleCate/updateCpsArticleCate [put]
func (cpsArticleCateApi *CpsArticleCateApi) UpdateCpsArticleCate(c *gin.Context) {
	var cpsArticleCate autocode.CpsArticleCate
	_ = c.ShouldBindJSON(&cpsArticleCate)
	if err := cpsArticleCateService.UpdateCpsArticleCate(cpsArticleCate); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCpsArticleCate 用id查询CpsArticleCate
// @Tags CpsArticleCate
// @Summary 用id查询CpsArticleCate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CpsArticleCate true "用id查询CpsArticleCate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cpsArticleCate/findCpsArticleCate [get]
func (cpsArticleCateApi *CpsArticleCateApi) FindCpsArticleCate(c *gin.Context) {
	var cpsArticleCate autocode.CpsArticleCate
	_ = c.ShouldBindQuery(&cpsArticleCate)
	if err, recpsArticleCate := cpsArticleCateService.GetCpsArticleCate(cpsArticleCate.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recpsArticleCate": recpsArticleCate}, c)
	}
}

// GetCpsArticleCateList 分页获取CpsArticleCate列表
// @Tags CpsArticleCate
// @Summary 分页获取CpsArticleCate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CpsArticleCateSearch true "分页获取CpsArticleCate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpsArticleCate/getCpsArticleCateList [get]
func (cpsArticleCateApi *CpsArticleCateApi) GetCpsArticleCateList(c *gin.Context) {
	var pageInfo autocodeReq.CpsArticleCateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cpsArticleCateService.GetCpsArticleCateInfoList(pageInfo); err != nil {
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

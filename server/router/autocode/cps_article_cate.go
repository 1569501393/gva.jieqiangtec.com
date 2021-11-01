package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CpsArticleCateRouter struct {
}

// InitCpsArticleCateRouter 初始化 CpsArticleCate 路由信息
func (s *CpsArticleCateRouter) InitCpsArticleCateRouter(Router *gin.RouterGroup) {
	cpsArticleCateRouter := Router.Group("cpsArticleCate").Use(middleware.OperationRecord())
	cpsArticleCateRouterWithoutRecord := Router.Group("cpsArticleCate")
	var cpsArticleCateApi = v1.ApiGroupApp.AutoCodeApiGroup.CpsArticleCateApi
	{
		cpsArticleCateRouter.POST("createCpsArticleCate", cpsArticleCateApi.CreateCpsArticleCate)   // 新建CpsArticleCate
		cpsArticleCateRouter.DELETE("deleteCpsArticleCate", cpsArticleCateApi.DeleteCpsArticleCate) // 删除CpsArticleCate
		cpsArticleCateRouter.DELETE("deleteCpsArticleCateByIds", cpsArticleCateApi.DeleteCpsArticleCateByIds) // 批量删除CpsArticleCate
		cpsArticleCateRouter.PUT("updateCpsArticleCate", cpsArticleCateApi.UpdateCpsArticleCate)    // 更新CpsArticleCate
	}
	{
		cpsArticleCateRouterWithoutRecord.GET("findCpsArticleCate", cpsArticleCateApi.FindCpsArticleCate)        // 根据ID获取CpsArticleCate
		cpsArticleCateRouterWithoutRecord.GET("getCpsArticleCateList", cpsArticleCateApi.GetCpsArticleCateList)  // 获取CpsArticleCate列表
	}
}

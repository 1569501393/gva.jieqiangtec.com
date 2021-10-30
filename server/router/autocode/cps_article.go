package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CpsArticleRouter struct {
}

// InitCpsArticleRouter 初始化 CpsArticle 路由信息
func (s *CpsArticleRouter) InitCpsArticleRouter(Router *gin.RouterGroup) {
	cpsArticleRouter := Router.Group("cpsArticle").Use(middleware.OperationRecord())
	cpsArticleRouterWithoutRecord := Router.Group("cpsArticle")
	var cpsArticleApi = v1.ApiGroupApp.AutoCodeApiGroup.CpsArticleApi
	{
		cpsArticleRouter.POST("createCpsArticle", cpsArticleApi.CreateCpsArticle)   // 新建CpsArticle
		cpsArticleRouter.DELETE("deleteCpsArticle", cpsArticleApi.DeleteCpsArticle) // 删除CpsArticle
		cpsArticleRouter.DELETE("deleteCpsArticleByIds", cpsArticleApi.DeleteCpsArticleByIds) // 批量删除CpsArticle
		cpsArticleRouter.PUT("updateCpsArticle", cpsArticleApi.UpdateCpsArticle)    // 更新CpsArticle
	}
	{
		cpsArticleRouterWithoutRecord.GET("findCpsArticle", cpsArticleApi.FindCpsArticle)        // 根据ID获取CpsArticle
		cpsArticleRouterWithoutRecord.GET("getCpsArticleList", cpsArticleApi.GetCpsArticleList)  // 获取CpsArticle列表
	}
}

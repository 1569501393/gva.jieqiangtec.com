package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CpsArticleService struct {
}

// CreateCpsArticle 创建CpsArticle记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleService *CpsArticleService) CreateCpsArticle(cpsArticle autocode.CpsArticle) (err error) {
	err = global.GVA_DB.Create(&cpsArticle).Error
	return err
}

// DeleteCpsArticle 删除CpsArticle记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleService *CpsArticleService)DeleteCpsArticle(cpsArticle autocode.CpsArticle) (err error) {
	err = global.GVA_DB.Delete(&cpsArticle).Error
	return err
}

// DeleteCpsArticleByIds 批量删除CpsArticle记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleService *CpsArticleService)DeleteCpsArticleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CpsArticle{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCpsArticle 更新CpsArticle记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleService *CpsArticleService)UpdateCpsArticle(cpsArticle autocode.CpsArticle) (err error) {
	err = global.GVA_DB.Save(&cpsArticle).Error
	return err
}

// GetCpsArticle 根据id获取CpsArticle记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleService *CpsArticleService)GetCpsArticle(id uint) (err error, cpsArticle autocode.CpsArticle) {
	err = global.GVA_DB.Where("id = ?", id).First(&cpsArticle).Error
	return
}

// GetCpsArticleInfoList 分页获取CpsArticle记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleService *CpsArticleService)GetCpsArticleInfoList(info autoCodeReq.CpsArticleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.CpsArticle{})
    var cpsArticles []autocode.CpsArticle
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&cpsArticles).Error
	return err, cpsArticles, total
}

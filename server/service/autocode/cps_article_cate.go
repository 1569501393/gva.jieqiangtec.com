package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CpsArticleCateService struct {
}

// CreateCpsArticleCate 创建CpsArticleCate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleCateService *CpsArticleCateService) CreateCpsArticleCate(cpsArticleCate autocode.CpsArticleCate) (err error) {
	err = global.GVA_DB.Create(&cpsArticleCate).Error
	return err
}

// DeleteCpsArticleCate 删除CpsArticleCate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleCateService *CpsArticleCateService)DeleteCpsArticleCate(cpsArticleCate autocode.CpsArticleCate) (err error) {
	err = global.GVA_DB.Delete(&cpsArticleCate).Error
	return err
}

// DeleteCpsArticleCateByIds 批量删除CpsArticleCate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleCateService *CpsArticleCateService)DeleteCpsArticleCateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CpsArticleCate{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCpsArticleCate 更新CpsArticleCate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleCateService *CpsArticleCateService)UpdateCpsArticleCate(cpsArticleCate autocode.CpsArticleCate) (err error) {
	err = global.GVA_DB.Save(&cpsArticleCate).Error
	return err
}

// GetCpsArticleCate 根据id获取CpsArticleCate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleCateService *CpsArticleCateService)GetCpsArticleCate(id uint) (err error, cpsArticleCate autocode.CpsArticleCate) {
	err = global.GVA_DB.Where("id = ?", id).First(&cpsArticleCate).Error
	return
}

// GetCpsArticleCateInfoList 分页获取CpsArticleCate记录
// Author [piexlmax](https://github.com/piexlmax)
func (cpsArticleCateService *CpsArticleCateService)GetCpsArticleCateInfoList(info autoCodeReq.CpsArticleCateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.CpsArticleCate{})
    var cpsArticleCates []autocode.CpsArticleCate
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Alias != "" {
        db = db.Where("`alias` = ?",info.Alias)
    }
    if info.ArticleNums != nil {
        db = db.Where("`article_nums` = ?",info.ArticleNums)
    }
    if info.Name != "" {
        db = db.Where("`name` = ?",info.Name)
    }
    if info.Pid != nil {
        db = db.Where("`pid` = ?",info.Pid)
    }
    if info.SeoTitle != "" {
        db = db.Where("`seo_title` = ?",info.SeoTitle)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&cpsArticleCates).Error
	return err, cpsArticleCates, total
}

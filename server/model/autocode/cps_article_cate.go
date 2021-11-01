// 自动生成模板CpsArticleCate
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CpsArticleCate 结构体
// 如果含有time.Time 请自行import time包
type CpsArticleCate struct {
      global.GVA_MODEL
      Alias  string `json:"alias" form:"alias" gorm:"column:alias;comment:;type:varchar(50);"`
      ArticleNums  *int `json:"articleNums" form:"articleNums" gorm:"column:article_nums;comment:;type:int"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:分类名;type:varchar(100);"`
      Pid  *int `json:"pid" form:"pid" gorm:"column:pid;comment:父id;type:smallint"`
      SeoDesc  string `json:"seoDesc" form:"seoDesc" gorm:"column:seo_desc;comment:;type:text;"`
      SeoKeys  string `json:"seoKeys" form:"seoKeys" gorm:"column:seo_keys;comment:;type:varchar(255);"`
      SeoTitle  string `json:"seoTitle" form:"seoTitle" gorm:"column:seo_title;comment:;type:varchar(255);"`
      SortOrder  *int `json:"sortOrder" form:"sortOrder" gorm:"column:sort_order;comment:;type:smallint"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:;type:tinyint"`
}


// TableName CpsArticleCate 表名
func (CpsArticleCate) TableName() string {
  return "cps_article_cate"
}


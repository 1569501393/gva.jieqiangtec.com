// 自动生成模板Article
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Article 结构体
// 如果含有time.Time 请自行import time包
type Article struct {
      global.GVA_MODEL
      Id  *int `json:"id" form:"id" gorm:"column:id;comment:自增id;type:int"`
}


// TableName Article 表名
func (Article) TableName() string {
  return "cps_article"
}


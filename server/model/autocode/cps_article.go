// 自动生成模板CpsArticle
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
      "time"
)

// CpsArticle 结构体
// 如果含有time.Time 请自行import time包
type CpsArticle struct {
      global.GVA_MODEL
      Abst  string `json:"abst" form:"abst" gorm:"column:abst;comment:;type:varchar(255);"`
      AddTime  *time.Time `json:"addTime" form:"addTime" gorm:"column:add_time;comment:;type:datetime"`
      Aid  string `json:"aid" form:"aid" gorm:"column:aid;comment:附件id;type:varchar(255);"`
      CateId  *bool `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:;type:tinyint"`
      DataState  *bool `json:"dataState" form:"dataState" gorm:"column:data_state;comment:数据状态：0删除，1正常;type:tinyint"`
      Img  string `json:"img" form:"img" gorm:"column:img;comment:;type:varchar(255);"`
      Info  string `json:"info" form:"info" gorm:"column:info;comment:信息;type:mediumtext;"`
      IsBest  *bool `json:"isBest" form:"isBest" gorm:"column:is_best;comment:;type:tinyint"`
      IsHot  *bool `json:"isHot" form:"isHot" gorm:"column:is_hot;comment:;type:tinyint"`
      Ordid  *bool `json:"ordid" form:"ordid" gorm:"column:ordid;comment:;type:tinyint"`
      Orig  string `json:"orig" form:"orig" gorm:"column:orig;comment:;type:varchar(255);"`
      PlatformId  *int `json:"platformId" form:"platformId" gorm:"column:platform_id;comment:分销机构（发布的时候可指定全部或者具体分行、子机构的人员能看到） 0全部;type:int"`
      SeoDesc  string `json:"seoDesc" form:"seoDesc" gorm:"column:seo_desc;comment:;type:varchar(1024);"`
      SeoKeys  string `json:"seoKeys" form:"seoKeys" gorm:"column:seo_keys;comment:;type:varchar(255);"`
      SeoTitle  string `json:"seoTitle" form:"seoTitle" gorm:"column:seo_title;comment:;type:varchar(255);"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:0-待审核 1-已审核;type:tinyint"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;type:varchar(255);"`
      Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:创建人;type:int"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:;type:varchar(255);"`
}


// TableName CpsArticle 表名
func (CpsArticle) TableName() string {
  return "cps_article"
}


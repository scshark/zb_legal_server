// 自动生成模板DocumentCategory
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type ZbDocumentCategory struct {
      gorm.Model
      Title  string `json:"title" form:"title" gorm:"column:title;comment:'分类名称';type:varchar(255)"`
      ParentId  int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:'父分类ID';type:int(10)"`
      CategoryLevel  int `json:"categoryLevel" form:"categoryLevel" gorm:"column:category_level;comment:'分类等级';type:int(10)"`
      Sort  int `json:"sort" form:"sort" gorm:"column:sort;comment:'排序';type:smallint(3)"` 
}


func (ZbDocumentCategory) TableName() string {
  return "zb_document_category"
}

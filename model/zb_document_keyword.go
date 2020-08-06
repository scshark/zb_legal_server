// 自动生成模板DocumentKeyword
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type ZbDocumentKeyword struct {
      gorm.Model
      Name  string `json:"name" form:"name" gorm:"column:name;comment:'名称';type:varchar(255)"`
      SearchNum  int `json:"searchNum" form:"searchNum" gorm:"column:search_num;comment:'搜索量';type:int(10)"`
      DocumentNum  int `json:"documentNum" form:"documentNum" gorm:"column:document_num;comment:'文书数量';type:int(10)"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:'状态 0:禁用 1:正常';type:smallint(1)"`
      Sort  int `json:"sort" form:"sort" gorm:"column:sort;comment:'排序 ';type:smallint(3)"` 
}


func (ZbDocumentKeyword) TableName() string {
  return "zb_document_keyword"
}

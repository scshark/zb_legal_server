// 自动生成模板Document
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type ZbDocumentCategoryRelation struct {
	gorm.Model
	DocumentId int `json:"documentId" form:"documentId" gorm:"column:document_id;comment:'文书id';type:int(10)"`
	CategoryId int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:'文书分类ID';type:int(10)"`
	Status     int `json:"status" form:"status" gorm:"column:status;comment:'状态 0：不可用 1：可用';type:smallint(1);default:1"`
	Sort       int `json:"sort" form:"sort" gorm:"column:sort;comment:'排序';type:smallint(3);default:0"`
}

func (ZbDocumentCategoryRelation) TableName() string {
	return "zb_document_category_relation"
}

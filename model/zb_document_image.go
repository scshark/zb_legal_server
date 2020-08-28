// 自动生成模板Document
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type ZbDocumentImage struct {
	gorm.Model
	DocumentId int    `json:"documentId" form:"documentId" gorm:"column:document_id;comment:'文书id';type:int(10)"`
	ImageUrl   string `json:"imageUrl" form:"imageUrl" gorm:"column:image_url;comment:'图片地址';type:varchar(255)"`
	Status     int    `json:"status" form:"status" gorm:"column:status;comment:'状态 0：不可用 1：可用';type:tinyint(1)"`
	Sort       int    `json:"sort" form:"sort" gorm:"column:sort;comment:'排序';type:smallint(3)"`
}

func (ZbDocumentImage) TableName() string {
	return "zb_document_image"
}

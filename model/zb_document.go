// 自动生成模板Document
package model

import (
	"github.com/jinzhu/gorm"
    "time"
)

// 如果含有time.Time 请自行import time包
type ZbDocument struct {
      gorm.Model
      Title  string `json:"title" form:"title" gorm:"column:title;comment:'标题';type:varchar(255)"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:'内容';type:text"`
      WordFileUrl  string `json:"wordFileUrl" form:"wordFileUrl" gorm:"column:word_file_url;comment:'word文档下载地址';type:varchar(255)"`
      PdfFileUrl  string `json:"pdfFileUrl" form:"pdfFileUrl" gorm:"column:pdf_file_url;comment:'pdf文档下载地址';type:varchar(255)"`
      BrowseNum  int `json:"browseNum" form:"browseNum" gorm:"column:browse_num;comment:'浏览量';type:int(10)"`
      DownloadNum  int `json:"downloadNum" form:"downloadNum" gorm:"column:download_num;comment:'下载量';type:int(10)"`
      BrowseVirtualNum  int `json:"browseVirtualNum" form:"browseVirtualNum" gorm:"column:browse_virtual_num;comment:'虚拟浏览量';type:int(10)"`
      DownloadVirtualNum  int `json:"downloadVirtualNum" form:"downloadVirtualNum" gorm:"column:download_virtual_num;comment:'虚拟下载量';type:int(10)"`
      CollectionNum  int `json:"collectionNum" form:"collectionNum" gorm:"column:collection_num;comment:'收藏数量';type:int(10)"`
      ShareNum  int `json:"shareNum" form:"shareNum" gorm:"column:share_num;comment:'分享数量';type:int(10)"`
      ReleasedAt  time.Time `json:"releasedAt" form:"releasedAt" gorm:"column:released_at;comment:'发布日期';type:timestamp;default:null"`
      RevisedAt  time.Time `json:"revisedAt" form:"revisedAt" gorm:"column:revised_at;comment:'最后修改日期';type:timestamp;default:null"`
      ClassId  []int `json:"classId" form:"classId" gorm:"-"`
}

func (ZbDocument) TableName() string {
  return "zb_document"
}
